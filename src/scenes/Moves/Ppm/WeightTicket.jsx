import React, { Component, Fragment } from 'react';
import { reduxForm, getFormValues, SubmissionError } from 'redux-form';
import { connect } from 'react-redux';
import { get, map, isEmpty } from 'lodash';
import PropTypes from 'prop-types';

import { ProgressTimeline, ProgressTimelineStep } from 'shared/ProgressTimeline';
import { SwaggerField } from 'shared/JsonSchemaForm/JsonSchemaField';
import RadioButton from 'shared/RadioButton';
import Checkbox from 'shared/Checkbox';
import Uploader from 'shared/Uploader';
import { createMoveDocument } from 'shared/Entities/modules/moveDocuments';
import Alert from 'shared/Alert';

import PPMPaymentRequestActionBtns from './PPMPaymentRequestActionBtns';
import WizardHeader from '../WizardHeader';
import './PPMPaymentRequest.css';

class WeightTicket extends Component {
  state = {
    vehicleType: '',
    additionalWeightTickets: 'Yes',
    weightTicketSubmissionError: false,
    missingEmptyWeightTicket: false,
    missingFullWeightTicket: false,
  };
  uploaders = {};

  isMissingWeightTickets = () => {
    if (isEmpty(this.uploaders)) {
      return true;
    }
    const uploadersKeys = Object.keys(this.uploaders);
    for (const key of uploadersKeys) {
      // eslint-disable-next-line security/detect-object-injection
      if (this.uploaders[key].isEmpty()) {
        return true;
      }
    }
    return false;
  };

  formIsIncomplete = () => {
    const { formValues } = this.props;
    const isMissingFormInput = !(
      formValues &&
      formValues.vehicle_nickname &&
      formValues.vehicle_options &&
      formValues.empty_weight &&
      formValues.full_weight &&
      formValues.weight_ticket_date
    );
    return isMissingFormInput && this.isMissingWeightTickets();
  };

  //  handleChange for vehicleType and additionalWeightTickets
  handleChange = (event, type) => {
    this.setState({ [type]: event.target.value });
  };

  handleCheckboxChange = event => {
    this.setState({
      [event.target.name]: event.target.checked,
    });
  };

  onAddFile = uploaderName => () => {
    this.setState({
      uploaderIsIdle: { ...this.state.uploaderIsIdle, [uploaderName]: false },
    });
  };

  onChange = uploaderName => uploaderIsIdle => {
    this.setState({
      uploaderIsIdle: { ...this.state.uploaderIsIdle, [uploaderName]: uploaderIsIdle },
    });
  };

  cancelHandler = formValues => {
    const { history } = this.props;
    history.push('/');
  };

  saveForLaterHandler = formValues => {
    const { history } = this.props;
    return this.saveAndAddHandler(formValues).then(() => {
      if (this.state.weightTicketSubmissionError === false) {
        history.push('/');
      }
    });
  };

  saveAndAddHandler = formValues => {
    const { moveId, currentPpm } = this.props;

    const moveDocumentSubmissions = [];
    const uploadersKeys = Object.keys(this.uploaders);
    for (const key of uploadersKeys) {
      // eslint-disable-next-line security/detect-object-injection
      let files = this.uploaders[key].getFiles();
      if (files.length > 0) {
        const uploadIds = map(files, 'id');
        const weightTicket = {
          moveId: moveId,
          personallyProcuredMoveId: currentPpm.id,
          uploadIds: uploadIds,
          title: key,
          moveDocumentType: 'WEIGHT_TICKET',
          notes: formValues.notes,
        };
        moveDocumentSubmissions.push(
          this.props.createMoveDocument(weightTicket).catch(() => {
            throw new SubmissionError({ _error: 'Error creating move document' });
          }),
        );
      }
    }
    return Promise.all(moveDocumentSubmissions)
      .then(() => {
        this.setState({ weightTicketSubmissionError: false });
        this.cleanup();
      })
      .catch(e => {
        this.setState({ weightTicketSubmissionError: true });
      });
  };

  cleanup = () => {
    const { reset } = this.props;
    const uploaders = this.uploaders;
    const uploadersKeys = Object.keys(this.uploaders);
    for (const key of uploadersKeys) {
      // eslint-disable-next-line security/detect-object-injection
      uploaders[key].clearFiles();
    }
    reset();
  };

  render() {
    const { additionalWeightTickets, vehicleType, missingEmptyWeightTicket, missingFullWeightTicket } = this.state;
    const { handleSubmit, submitting, schema } = this.props;
    const nextBtnLabel = additionalWeightTickets === 'Yes' ? 'Save & Add Another' : 'Save & Continue';
    const uploadEmptyTicketLabel =
      '<span class="uploader-label">Drag & drop or <span class="filepond--label-action">click to upload empty weight ticket</span></span>';
    const uploadFullTicketLabel =
      '<span class="uploader-label">Drag & drop or <span class="filepond--label-action">click to upload full weight ticket</span></span>';
    return (
      <Fragment>
        <WizardHeader
          title="Weight tickets"
          right={
            <ProgressTimeline>
              <ProgressTimelineStep name="Weight" current />
              <ProgressTimelineStep name="Expenses" />
              <ProgressTimelineStep name="Review" />
            </ProgressTimeline>
          }
        />
        <form>
          {this.state.weightTicketSubmissionError && (
            <div className="usa-grid">
              <div className="usa-width-one-whole error-message">
                <Alert type="error" heading="An error occurred">
                  Something went wrong contacting the server.
                </Alert>
              </div>
            </div>
          )}
          <div className="usa-grid">
            <SwaggerField
              fieldName="vehicle_options"
              swagger={schema}
              onChange={event => this.handleChange(event, 'vehicleType')}
              value={vehicleType}
              required
            />
            <SwaggerField fieldName="vehicle_nickname" swagger={schema} required />
            {(vehicleType === 'CAR' || vehicleType === 'BOX_TRUCK') && (
              <>
                <div className="dashed-divider" />

                <div className="usa-grid-full" style={{ marginTop: '1em' }}>
                  <div className="usa-width-one-third">
                    <strong className="input-header">Empty Weight</strong>
                    <SwaggerField
                      className="short-field"
                      fieldName="empty_weight"
                      swagger={schema}
                      hideLabel
                      required
                    />{' '}
                    lbs
                  </div>
                  <div className="usa-width-two-thirds uploader-wrapper">
                    <Uploader
                      options={{ labelIdle: uploadEmptyTicketLabel }}
                      onRef={ref => (this.uploaders['empty_weight'] = ref)}
                      onChange={this.onChange('empty_weight')}
                      onAddFile={this.onAddFile('empty_weight')}
                    />
                    <Checkbox
                      label="I'm missing this weight ticket"
                      name="missingEmptyWeightTicket"
                      checked={missingEmptyWeightTicket}
                      onChange={this.handleCheckboxChange}
                      normalizeLabel
                    />
                    {missingEmptyWeightTicket && (
                      <Alert type="warning">
                        Contact your local Transportation Office (PPPO) to let them know you’re missing this weight
                        ticket. For now, keep going and enter the info you do have.
                      </Alert>
                    )}
                  </div>
                </div>
                <div className="usa-grid-full" style={{ marginTop: '1em' }}>
                  <div className="usa-width-one-third">
                    <strong className="input-header">Full Weight</strong>
                    <label className="full-weight-label">Full weight at destination</label>
                    <SwaggerField
                      className="short-field"
                      fieldName="full_weight"
                      swagger={schema}
                      hideLabel
                      required
                    />{' '}
                    lbs
                  </div>

                  <div className="usa-width-two-thirds uploader-wrapper">
                    <Uploader
                      options={{ labelIdle: uploadFullTicketLabel }}
                      onRef={ref => (this.uploaders['full_weight'] = ref)}
                      onChange={this.onChange('full_weight')}
                      onAddFile={this.onAddFile('full_weight')}
                    />
                    <Checkbox
                      label="I'm missing this weight ticket"
                      name="missingFullWeightTicket"
                      checked={missingFullWeightTicket}
                      onChange={this.handleCheckboxChange}
                      normalizeLabel
                    />
                    {missingFullWeightTicket && (
                      <Alert type="warning">
                        Contact your local Transportation Office (PPPO) to let them know you’re missing this weight
                        ticket. For now, keep going and enter the info you do have.
                      </Alert>
                    )}
                  </div>
                </div>

                <SwaggerField fieldName="weight_ticket_date" swagger={schema} required />
                <div className="dashed-divider" />

                <div className="radio-group-wrapper">
                  <p className="radio-group-header">Do you have more weight tickets for another vehicle or trip?</p>
                  <RadioButton
                    inputClassName="inline_radio"
                    labelClassName="inline_radio"
                    label="Yes"
                    value="Yes"
                    name="additional_weight_ticket"
                    checked={additionalWeightTickets === 'Yes'}
                    onChange={event => this.handleChange(event, 'additionalWeightTickets')}
                  />

                  <RadioButton
                    inputClassName="inline_radio"
                    labelClassName="inline_radio"
                    label="No"
                    value="No"
                    name="additional_weight_ticket"
                    checked={additionalWeightTickets === 'No'}
                    onChange={event => this.handleChange(event, 'additionalWeightTickets')}
                  />
                </div>
              </>
            )}

            {/* TODO: change onclick handler to go to next page in flow */}
            <PPMPaymentRequestActionBtns
              nextBtnLabel={nextBtnLabel}
              submitButtonsAreDisabled={this.formIsIncomplete()}
              submitting={submitting}
              cancelHandler={this.cancelHandler}
              saveForLaterHandler={handleSubmit(this.saveForLaterHandler)}
              saveAndAddHandler={handleSubmit(this.saveAndAddHandler)}
              displaySaveForLater={true}
            />
          </div>
        </form>
      </Fragment>
    );
  }
}

const formName = 'weight_ticket_wizard';
WeightTicket = reduxForm({
  form: formName,
  enableReinitialize: true,
  keepDirtyOnReinitialize: true,
})(WeightTicket);

WeightTicket.propTypes = {
  schema: PropTypes.object.isRequired,
};

function mapStateToProps(state, ownProps) {
  const moveId = ownProps.match.params.moveId;
  return {
    moveId: moveId,
    formValues: getFormValues(formName)(state),
    genericMoveDocSchema: get(state, 'swaggerInternal.spec.definitions.CreateGenericMoveDocumentPayload', {}),
    moveDocSchema: get(state, 'swaggerInternal.spec.definitions.MoveDocumentPayload', {}),
    schema: get(state, 'swaggerInternal.spec.definitions.WeightTicketPayload', {}),
    currentPpm: get(state, 'ppm.currentPpm'),
  };
}
const mapDispatchToProps = {
  createMoveDocument,
};

export default connect(mapStateToProps, mapDispatchToProps)(WeightTicket);
