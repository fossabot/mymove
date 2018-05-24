import React, { Component } from 'react';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { get } from 'lodash';

import { push } from 'react-router-redux';
import { reduxForm, FormSection } from 'redux-form';
import Alert from 'shared/Alert'; // eslint-disable-line

import { SwaggerField } from 'shared/JsonSchemaForm/JsonSchemaField';

import { updateServiceMember } from 'scenes/ServiceMembers/ducks';

import 'scenes/ServiceMembers/ServiceMembers.css';
import './Review.css';

const editContactFormName = 'edit_contact_info';

let EditContactForm = props => {
  const {
    onCancel,
    serviceMemberSchema,
    addressSchema,
    handleSubmit,
    submitting,
    valid,
  } = props;
  return (
    <form className="service_member_contact_info" onSubmit={handleSubmit}>
      <FormSection name="serviceMember">
        <h2>Edit Contact Info</h2>
        <SwaggerField
          fieldName="telephone"
          swagger={serviceMemberSchema}
          required
        />
        <SwaggerField
          fieldName="secondary_telephone"
          swagger={serviceMemberSchema}
        />
        <SwaggerField
          fieldName="personal_email"
          swagger={serviceMemberSchema}
          required
        />
        <fieldset key="contact_preferences">
          <legend htmlFor="contact_preferences">
            Preferred contact method(s) during your move:
          </legend>
          <SwaggerField
            fieldName="phone_is_preferred"
            swagger={serviceMemberSchema}
          />
          <SwaggerField
            fieldName="text_message_is_preferred"
            swagger={serviceMemberSchema}
          />
          <SwaggerField
            fieldName="email_is_preferred"
            swagger={serviceMemberSchema}
          />
        </fieldset>
      </FormSection>
      <hr className="spacer" />

      <FormSection name="resAddress">
        <h3>Current Residence Address</h3>
        <SwaggerField
          fieldName="street_address_1"
          swagger={addressSchema}
          required
        />
        <SwaggerField fieldName="street_address_2" swagger={addressSchema} />
        <SwaggerField fieldName="city" swagger={addressSchema} required />
        <SwaggerField fieldName="state" swagger={addressSchema} required />
        <SwaggerField
          fieldName="postal_code"
          swagger={addressSchema}
          required
        />
      </FormSection>
      <hr className="spacer" />
      <FormSection name="backupAddress">
        <h3>Backup Mailing Address</h3>
        <SwaggerField
          fieldName="street_address_1"
          swagger={addressSchema}
          required
        />
        <SwaggerField fieldName="street_address_2" swagger={addressSchema} />
        <SwaggerField fieldName="city" swagger={addressSchema} required />
        <SwaggerField fieldName="state" swagger={addressSchema} required />
        <SwaggerField
          fieldName="postal_code"
          swagger={addressSchema}
          required
        />
      </FormSection>
      <button type="submit" disabled={submitting || !valid}>
        Save
      </button>
      <button type="button" disabled={submitting} onClick={onCancel}>
        Cancel
      </button>
    </form>
  );
};

const validateEditContactFormBools = fields => {
  return (values, form) => {
    let errors = {};
    let prefSelected = false;
    fields.forEach(fieldName => {
      if (Boolean(get(values, fieldName))) {
        prefSelected = true;
      }
    });
    if (!prefSelected) {
      let valueSection = fields[0].split('.')[0];
      let field = fields[0].split('.')[1];
      var errorMsg = {
        [field]: 'Please select a preferred method of contact.',
      };
      var newError = { [valueSection]: errorMsg };
      return newError;
    }
    return errors;
  };
};

EditContactForm = reduxForm({
  form: editContactFormName,
  validate: validateEditContactFormBools([
    'serviceMember.phone_is_preferred',
    'serviceMember.text_message_is_preferred',
    'serviceMember.email_is_preferred',
  ]),
})(EditContactForm);

class EditContact extends Component {
  returnToReview = () => {
    const reviewAddress = `/moves/${this.props.match.params.moveId}/review`;
    this.props.push(reviewAddress);
  };

  updateContact = fieldValues => {
    let serviceMember = fieldValues.serviceMember;
    serviceMember.residential_address = fieldValues.resAddress;
    serviceMember.backup_mailing_address = fieldValues.backupAddress;
    return this.props.updateServiceMember(serviceMember).then(() => {
      // This promise resolves regardless of error.
      if (!this.props.hasSubmitError) {
        this.returnToReview();
      } else {
        window.scrollTo(0, 0);
      }
    });
  };
  render() {
    const {
      error,
      serviceMemberSchema,
      addressSchema,
      serviceMember,
    } = this.props;
    let initialValues = null;
    if (
      serviceMember &&
      get(serviceMember, 'residential_address') &&
      get(serviceMember, 'backup_mailing_address')
    )
      initialValues = {
        serviceMember: serviceMember,
        resAddress: serviceMember.residential_address,
        backupAddress: serviceMember.backup_mailing_address,
      };
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
          <EditContactForm
            initialValues={initialValues}
            serviceMemberSchema={serviceMemberSchema}
            addressSchema={addressSchema}
            onSubmit={this.updateContact}
            onCancel={this.returnToReview}
          />
        </div>
      </div>
    );
  }
}

function mapStateToProps(state) {
  let serviceMember = get(
    state,
    'loggedInUser.loggedInUser.service_member',
    {},
  );
  return {
    serviceMember: serviceMember,
    move: get(state, 'moves.currentMove'),
    error: get(state, 'serviceMember.error'),
    hasSubmitError: get(state, 'serviceMember.hasSubmitError'),
    serviceMemberSchema: get(
      state,
      'swagger.spec.definitions.CreateServiceMemberPayload',
      {},
    ),
    addressSchema: get(state, 'swagger.spec.definitions.Address', {}),
  };
}

function mapDispatchToProps(dispatch) {
  return bindActionCreators({ push, updateServiceMember }, dispatch);
}

export default connect(mapStateToProps, mapDispatchToProps)(EditContact);
