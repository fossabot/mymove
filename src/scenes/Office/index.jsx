import React, { Component } from 'react';
import { Redirect, Route, Switch } from 'react-router-dom';
import { ConnectedRouter } from 'react-router-redux';
import { history } from 'shared/store';
import { connect } from 'react-redux';
import { bindActionCreators } from 'redux';

import QueueHeader from 'shared/Header/Office';
import QueueList from './QueueList';
import QueueTable from './QueueTable';
import MoveInfo from './MoveInfo';
import OrdersInfo from './OrdersInfo';
import DocumentViewer from './DocumentViewer';
import { getCurrentUserInfo } from 'shared/Data/users';
import { loadInternalSchema, loadPublicSchema } from 'shared/Swagger/ducks';
import { detectIE11, no_op } from 'shared/utils';
import LogoutOnInactivity from 'shared/User/LogoutOnInactivity';
import PrivateRoute from 'shared/User/PrivateRoute';
import ScratchPad from 'shared/ScratchPad';
import { isProduction } from 'shared/constants';
import SomethingWentWrong from 'shared/SomethingWentWrong';
import { RetrieveMovesForOffice } from './api';

import './office.scss';

export class Queues extends Component {
  render() {
    return (
      <div className="usa-grid grid-wide queue-columns">
        <div className="queue-menu-column">
          <QueueList />
        </div>
        <div className="queue-list-column">
          <QueueTable queueType={this.props.match.params.queueType} retrieveMoves={RetrieveMovesForOffice} />
        </div>
      </div>
    );
  }
}

export class RenderWithHeader extends Component {
  render() {
    const Tag = detectIE11() ? 'div' : 'main';
    const Component = this.props.component;
    return (
      <>
        <QueueHeader />
        <Tag role="main" className="site__content">
          <Component {...this.props} />
        </Tag>
      </>
    );
  }
}

export class RenderWithoutHeader extends Component {
  render() {
    const Tag = detectIE11() ? 'div' : 'main';
    const Component = this.props.component;
    return (
      <>
        <Tag role="main" className="site__content">
          <Component {...this.props} />
        </Tag>
      </>
    );
  }
}

export class OfficeWrapper extends Component {
  state = { hasError: false };

  componentDidMount() {
    document.title = 'Transcom PPP: Office';
    this.props.loadInternalSchema();
    this.props.loadPublicSchema();
    this.props.getCurrentUserInfo();
  }

  componentDidCatch(error, info) {
    this.setState({
      hasError: true,
      error,
      info,
    });
  }

  render() {
    const ConditionalWrap = ({ condition, wrap, children }) => (condition ? wrap(children) : <>{children}</>);
    const Tag = detectIE11() ? 'div' : 'main';
    const userIsLoggedOff = this.props.userHasErrored;

    return (
      <ConnectedRouter history={history}>
        <div className="Office site">
          {userIsLoggedOff && <QueueHeader />}
          <ConditionalWrap
            condition={userIsLoggedOff}
            wrap={children => (
              <Tag role="main" className="site__content">
                {children}
              </Tag>
            )}
          >
            <LogoutOnInactivity />
            {this.state.hasError && <SomethingWentWrong error={this.state.error} info={this.state.info} />}
            {!this.state.hasError && (
              <Switch>
                <Route
                  exact
                  path="/"
                  component={({ location }) => (
                    <Redirect
                      from="/"
                      to={{
                        ...location,
                        pathname: '/queues/new',
                      }}
                    />
                  )}
                />
                <PrivateRoute
                  path="/queues/:queueType/moves/:moveId"
                  component={props => <RenderWithHeader component={MoveInfo} {...props} />}
                />
                <PrivateRoute
                  path="/queues/:queueType"
                  component={props => <RenderWithHeader component={Queues} {...props} />}
                />
                <PrivateRoute
                  path="/moves/:moveId/orders"
                  component={props => <RenderWithoutHeader component={OrdersInfo} {...props} />}
                />
                <PrivateRoute
                  path="/moves/:moveId/documents/:moveDocumentId?"
                  component={props => <RenderWithoutHeader component={DocumentViewer} {...props} />}
                />
                {!isProduction && (
                  <PrivateRoute
                    path="/playground"
                    component={props => <RenderWithHeader component={ScratchPad} {...props} />}
                  />
                )}
              </Switch>
            )}
          </ConditionalWrap>
        </div>
      </ConnectedRouter>
    );
  }
}

OfficeWrapper.defaultProps = {
  loadInternalSchema: no_op,
  loadPublicSchema: no_op,
};

const mapStateToProps = state => ({
  swaggerError: state.swaggerInternal.hasErrored,
  userHasErrored: state.user.hasErrored,
});

const mapDispatchToProps = dispatch =>
  bindActionCreators({ loadInternalSchema, loadPublicSchema, getCurrentUserInfo }, dispatch);

export default connect(mapStateToProps, mapDispatchToProps)(OfficeWrapper);
