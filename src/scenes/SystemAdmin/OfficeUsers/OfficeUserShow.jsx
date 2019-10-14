import React from 'react';
import { Show, SimpleShowLayout, TextField, BooleanField, DateField } from 'react-admin';

const OfficeUserShowTitle = ({ record }) => {
  return <span>{`${record.first_name} ${record.last_name}`}</span>;
};

const OfficeUserShow = props => {
  return (
    <Show {...props} title={<OfficeUserShowTitle />}>
      <SimpleShowLayout>
        <TextField source="id" />
        <TextField source="email" />
        <TextField source="first_name" />
        <TextField source="middle_initials" />
        <TextField source="last_name" />
        <TextField source="telephone" />
        <BooleanField source="disabled" label="Deactivated" />
        <DateField source="created_at" showTime />
        <DateField source="updated_at" showTime />
      </SimpleShowLayout>
    </Show>
  );
};

export default OfficeUserShow;