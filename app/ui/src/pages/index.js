import './index.scss';

import React from 'react';
import txt from 'lang';
import Invoices from './invoices/invoices';
import { connect } from 'react-redux';
import { Link } from 'react-router-dom';
import { withRouter } from 'react-router';

import {
  FocusStyleManager,
  Tabs2 as Tabs,
  Tab2 as Tab,
} from '@blueprintjs/core';

FocusStyleManager.onlyShowFocusOnTabs();

const AppComponent = (props) => {
  console.log(props);

  return (
    <header>
      <nav className="pt-navbar pt-dark">
        <div className="pt-navbar-group pt-align-left">
          <Link to="/"><div className="pt-navbar-heading">contabi</div></Link>
          <Link to="/invoices"><button className="pt-button pt-active pt-minimal">{txt('Sales')}</button></Link>
          <button className="pt-button pt-minimal">{txt('Expenses')}</button>
          <button className="pt-button pt-minimal">{txt('Bank')}</button>
        </div>
        <div className="pt-navbar-group pt-align-right">
          <Link to="/login"><button className="pt-button pt-minimal pt-icon-user" /></Link>
          <button className="pt-button pt-minimal pt-icon-cog" />
        </div>
      </nav>
      <nav className="subnav">
        <Tabs animate renderActiveTabPanelOnly id="navbar" onChange={console.log}>
          <Tab id="Home" title={txt('Invoices')} />
          <Tab id="Files" title={txt('Clients')} />
        </Tabs>
      </nav>
    </header>
  );
};

const mapStateToProps = (state) => ({state});

const App = withRouter(connect(mapStateToProps, null)(AppComponent));

export { App };
export { default as Login } from './login/login';
export { default as Invoices } from './invoices/invoices';
