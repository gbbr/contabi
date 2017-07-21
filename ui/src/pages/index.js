import './index.scss';

import React from 'react';
import txt from 'lang';
import Invoices from './invoices/invoices'

import {
  FocusStyleManager,
  Tabs2 as Tabs,
  Tab2 as Tab,
} from '@blueprintjs/core';

FocusStyleManager.onlyShowFocusOnTabs();

const App = () => (
  <header>
    <nav className="pt-navbar pt-dark">
      <div className="pt-navbar-group pt-align-left">
        <div className="pt-navbar-heading">contabi</div>
        <button className="pt-button pt-active pt-minimal">{txt('Sales')}</button>
        <button className="pt-button pt-minimal">{txt('Expenses')}</button>
        <button className="pt-button pt-minimal">{txt('Bank')}</button>
      </div>
      <div className="pt-navbar-group pt-align-right">
        <button className="pt-button pt-minimal pt-icon-user"></button>
        <button className="pt-button pt-minimal pt-icon-cog"></button>
      </div>
    </nav>
    <nav className="subnav">
      <Tabs animate renderActiveTabPanelOnly id="navbar" onChange={console.log}>
          <Tab id="Home" title={txt('Invoices')} panel={<Invoices />} />
          <Tab id="Files" title={txt('Clients')} panel={<div className="page">B</div>} />
      </Tabs>
    </nav>
  </header>
);

export { App };
export { default as Login } from './login/login';
export { default as Invoices } from './invoices/invoices';
