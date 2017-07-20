import 'styles/index.scss';

import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import txt from 'lang';

import {
  Intent,
  AnchorButton,
  FocusStyleManager,
  Tabs2 as Tabs,
  Tab2 as Tab,
} from '@blueprintjs/core';

FocusStyleManager.onlyShowFocusOnTabs();

const MyComponent = () => (
  <div>
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
    <div className="subnav">
      <Tabs animate renderActiveTabPanelOnly id="navbar" onChange={console.log}>
          <Tab id="Home" title={txt('Invoices')} panel={<div>A</div>} />
          <Tab id="Files" title={txt('Clients')} panel={<div>B</div>} />
      </Tabs>
    </div>
  </div>
);

document.addEventListener("DOMContentLoaded", () => {
  ReactDOM.render(<MyComponent />, document.getElementById('app-root'));
});
