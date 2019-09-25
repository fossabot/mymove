select * into temp tempsit from storage_in_transits;
select * into temp tempshipment from shipments;
select * into temp tempsli from shipment_line_items;
-- Select all service members, orders and moves that are not in personally_procured_moves table and HHG move type
-- and make sure that we exclude service members that have combo moves or PPMs
select sm.id as service_member_id, o.id as order_id, m.id as move_id, o.uploaded_orders_id, sm.residential_address_id, sm.backup_mailing_address_id into temp tempsom
	from service_members sm
		inner join orders o on sm.id = o.service_member_id
		inner join moves m on m.orders_id = o.id
		WHERE m.selected_move_type = 'HHG'
			and m.id NOT IN (SELECT move_id FROM personally_procured_moves)
			and sm.id NOT IN (select sm.id as service_member_id from service_members sm
					inner join orders o on sm.id = o.service_member_id
					inner join moves m on m.orders_id = o.id
					WHERE m.selected_move_type <> 'HHG');
select * into temp tempdc from distance_calculations where id IN (select shipping_distance_id from tempshipment);

DROP TABLE IF EXISTS shipment_line_items;
DROP TABLE IF EXISTS shipment_line_item_dimensions;
DROP TABLE IF EXISTS shipment_offers;
DROP TABLE IF EXISTS service_agents;
DROP TABLE IF EXISTS shipment_recalculates;
DROP TABLE IF EXISTS shipment_recalculate_logs;
DROP TABLE IF EXISTS storage_in_transits;
DROP TABLE IF EXISTS storage_in_transit_number_trackers;
DROP TABLE IF EXISTS gbl_number_trackers;
DROP TABLE IF EXISTS blackout_dates;

ALTER TABLE move_documents DROP COLUMN IF EXISTS shipment_id;
ALTER TABLE invoices DROP COLUMN IF EXISTS shipment_id;
ALTER TABLE signed_certifications DROP COLUMN IF EXISTS shipment_id;

-- remove tsp_users table and disable associated tsp_users in users table
-- make sure that users that are also office users are not disabled
UPDATE USERS SET disabled = TRUE
	WHERE id IN (select user_id from tsp_users where user_id is not null)
		and id NOT IN (select user_id from office_users where user_id is not null)
		and id NOT IN (select user_id from admin_users where user_id is not null);
DROP TABLE tsp_users;

-- removing hhg moves
-- documents
DELETE FROM weight_ticket_set_documents WHERE move_document_id IN (SELECT md.id FROM move_documents md INNER JOIN tempsom m ON m.move_id = md.move_id);
DELETE FROM moving_expense_documents WHERE move_document_id IN (SELECT md.id FROM move_documents md INNER JOIN tempsom m ON m.move_id = md.move_id);

-- TODO: need to refactor this part
DELETE FROM move_documents WHERE move_id IN (SELECT move_id FROM tempsom);
DELETE FROM signed_certifications WHERE move_id IN (SELECT move_id FROM tempsom);

-- finally dropping the shipments
DROP TABLE IF EXISTS shipments;

-- Dropping moves that are select HHG
-- make sure that the moves don't have PPMs previously
DELETE FROM moves WHERE id in (select move_id from tempsom);

-- service members
DELETE FROM access_codes WHERE service_member_id IN (select service_member_id from tempsom);
DELETE FROM backup_contacts WHERE service_member_id IN (select service_member_id from tempsom);
DELETE FROM orders WHERE id IN (select order_id from tempsom);

-- delete service member order document
-- might be some data left from the service member so delete based on service member id
DELETE FROM uploads WHERE document_id IN (select id from documents where service_member_id in (select service_member_id from tempsom));
DELETE FROM documents WHERE service_member_id IN (select service_member_id from tempsom);

-- delete notifications before service member
DELETE FROM notifications WHERE service_member_id IN (select service_member_id from tempsom);
-- finally delete service members
DELETE FROM service_members WHERE id IN (select service_member_id from tempsom);

-- delete distance calcs
DELETE FROM distance_calculations where id IN (select storage_in_transit_distance_id from tempsit);
DELETE FROM distance_calculations where id IN (select shipping_distance_id from tempshipment);