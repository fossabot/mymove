import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { get } from 'lodash';

import { push } from 'react-router-redux';
import { reduxForm, Field } from 'redux-form';

import Alert from 'shared/Alert'; // eslint-disable-line
import { SwaggerField } from 'shared/JsonSchemaForm/JsonSchemaField';
import DutyStationSearchBox from 'scenes/ServiceMembers/DutyStationSearchBox';
import YesNoBoolean from 'shared/Inputs/YesNoBoolean';
import Uploader from 'shared/Uploader';
import UploadsTable from 'shared/Uploader/UploadsTable';

import {
  updateOrders,
  deleteUpload,
  addUploads,
  showCurrentOrders,
} from 'scenes/Orders/ducks';

import './Review.css';
import profileImage from './images/profile.png';

const editOrdersFormName = 'edit_orders';

let EditOrdersForm = props => {
  const {
    onCancel,
    onDelete,
    onUpload,
    schema,
    handleSubmit,
    submitting,
    valid,
    initialValues,
    existingUploads,
    newUploads,
  } = props;
  const hasUploads = newUploads.length || existingUploads.length;
  return (
    <form onSubmit={handleSubmit}>
      <img src={profileImage} alt="" /> Orders
      <hr />
      <h3 className="sm-heading">Edit Orders:</h3>
      <SwaggerField fieldName="orders_type" swagger={schema} required />
      <SwaggerField fieldName="issue_date" swagger={schema} required />
      <SwaggerField fieldName="report_by_date" swagger={schema} required />
      <SwaggerField
        fieldName="has_dependents"
        swagger={schema}
        component={YesNoBoolean}
      />
      <br />
      <Field name="new_duty_station" component={DutyStationSearchBox} />
      <p>Uploads:</p>
      {Boolean(existingUploads.length) && (
        <UploadsTable uploads={existingUploads} onDelete={onDelete} />
      )}
      {Boolean(get(initialValues, 'uploaded_orders')) && (
        <Uploader
          document={initialValues.uploaded_orders}
          onChange={onUpload}
        />
      )}
      <button type="submit" disabled={submitting || !valid || !hasUploads}>
        Save
      </button>
      <button type="button" disabled={submitting} onClick={onCancel}>
        Cancel
      </button>
    </form>
  );
};

EditOrdersForm = reduxForm({
  form: editOrdersFormName,
})(EditOrdersForm);

class EditOrders extends Component {
  constructor(props) {
    super(props);

    this.state = {
      newUploads: [],
    };
  }

  returnToReview = () => {
    const reviewAddress = `/moves/${this.props.match.params.moveId}/review`;
    this.props.push(reviewAddress);
  };

  componentDidUpdate = prevProps => {
    // Once service member loads, load the backup contact.
    if (this.props.serviceMember && !prevProps.serviceMember) {
      this.props.showCurrentOrders(this.props.serviceMember.id);
    }
  };

  deleteFile = (e, uploadId) => {
    e.preventDefault();
    this.props.deleteUpload(uploadId);
  };

  handleNewUpload = uploads => {
    this.setState({ newUploads: uploads });
  };

  updateOrders = fieldValues => {
    fieldValues.new_duty_station_id = fieldValues.new_duty_station.id;
    this.props.addUploads(this.state.newUploads);
    return this.props.updateOrders(fieldValues.id, fieldValues).then(() => {
      // This promise resolves regardless of error.
      if (!this.props.hasSubmitError) {
        this.returnToReview();
      } else {
        window.scrollTo(0, 0);
      }
    });
  };

  render() {
    const { error, schema, currentOrders, existingUploads } = this.props;
    return (
      <div className="usa-grid">
        {error && (
          <div className="usa-width-one-whole error-message">
            <Alert type="error" heading="An error occurred">
              {error.message}
            </Alert>
          </div>
        )}
        <div className="usa-width-one-whole">
          <EditOrdersForm
            initialValues={currentOrders}
            onSubmit={this.updateOrders}
            onCancel={this.returnToReview}
            schema={schema}
            existingUploads={existingUploads}
            newUploads={this.state.newUploads}
            onUpload={this.handleNewUpload}
            onDelete={this.deleteFile}
          />
        </div>
      </div>
    );
  }
}

function mapStateToProps(state) {
  const props = {
    serviceMember: get(state, 'loggedInUser.loggedInUser.service_member'),
    error: get(state, 'orders.error'),
    hasSubmitError: get(state, 'orders.hasSubmitError'),

    schema: get(state, 'swagger.spec.definitions.CreateUpdateOrders', {}),
    formData: state.form[editOrdersFormName],
    currentOrders: get(state, 'orders.currentOrders'),
    existingUploads: get(
      state,
      'orders.currentOrders.uploaded_orders.uploads',
      [],
    ),
  };
  return props;
}

function mapDispatchToProps(dispatch) {
  return bindActionCreators(
    { push, updateOrders, addUploads, deleteUpload, showCurrentOrders },
    dispatch,
  );
}

export default connect(mapStateToProps, mapDispatchToProps)(EditOrders);
