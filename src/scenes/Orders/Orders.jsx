import { get } from 'lodash';
import PropTypes from 'prop-types';
import React, { Component } from 'react';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

import { Field } from 'redux-form';

import { createOrders, updateOrders } from './ducks';
import { reduxifyWizardForm } from 'shared/WizardPage/Form';
import DutyStationSearchBox from 'scenes/ServiceMembers/DutyStationSearchBox';
import YesNoBoolean from 'shared/Inputs/YesNoBoolean';
import { SwaggerField } from 'shared/JsonSchemaForm/JsonSchemaField';
import { validateAdditionalFields } from 'shared/JsonSchemaForm';

import './Orders.css';

const validateOrdersForm = validateAdditionalFields(['new_duty_station']);

const formName = 'orders_info';
const OrdersWizardForm = reduxifyWizardForm(formName, validateOrdersForm);

export class Orders extends Component {
  handleSubmit = () => {
    const pendingValues = Object.assign({}, this.props.formData.values);
    // Update if orders object already extant
    if (pendingValues) {
      pendingValues['service_member_id'] = this.props.serviceMemberId;
      pendingValues['new_duty_station_id'] = pendingValues.new_duty_station.id;
      pendingValues['has_dependents'] = pendingValues.has_dependents || false;
      if (this.props.currentOrders) {
        this.props.updateOrders(this.props.currentOrders.id, pendingValues);
      } else {
        this.props.createOrders(pendingValues);
      }
    }
  };

  render() {
    const {
      pages,
      pageKey,
      error,
      currentOrders,
      serviceMemberId,
      hasSubmitSuccess,
    } = this.props;
    // initialValues has to be null until there are values from the action since only the first values are taken
    const initialValues = currentOrders ? currentOrders : null;
    return (
      <OrdersWizardForm
        handleSubmit={this.handleSubmit}
        className={formName}
        pageList={pages}
        pageKey={pageKey}
        hasSucceeded={hasSubmitSuccess}
        serverError={error}
        initialValues={initialValues}
        additionalParams={{ serviceMemberId }}
      >
        <h1 className="sm-heading">Tell Us About Your Move Orders</h1>
        <SwaggerField
          fieldName="orders_type"
          swagger={this.props.schema}
          required
        />
        <SwaggerField
          fieldName="issue_date"
          swagger={this.props.schema}
          required
        />
        <SwaggerField
          fieldName="report_by_date"
          swagger={this.props.schema}
          required
        />
        <SwaggerField
          fieldName="has_dependents"
          swagger={this.props.schema}
          component={YesNoBoolean}
        />
        <Field name="new_duty_station" component={DutyStationSearchBox} />
      </OrdersWizardForm>
    );
  }
}
Orders.propTypes = {
  schema: PropTypes.object.isRequired,
  updateOrders: PropTypes.func.isRequired,
  currentOrders: PropTypes.object,
  error: PropTypes.object,
  hasSubmitSuccess: PropTypes.bool.isRequired,
};

function mapDispatchToProps(dispatch) {
  return bindActionCreators(
    {
      updateOrders,
      createOrders,
    },
    dispatch,
  );
}

function mapStateToProps(state) {
  const props = {
    serviceMemberId: get(state, 'serviceMember.currentServiceMember.id'),
    schema: get(state, 'swagger.spec.definitions.CreateUpdateOrders', {}),
    formData: state.form[formName],
    ...state.orders.currentOrders,
    hasSubmitSuccess: state.moves.hasSubmitSuccess,
  };
  return props;
}
export default connect(mapStateToProps, mapDispatchToProps)(Orders);
