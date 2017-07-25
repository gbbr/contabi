import './login.scss';

import React from 'react';
import txt from 'lang';
import { Dialog, Checkbox, Button, Intent } from '@blueprintjs/core';

const Login = () =>
  <Dialog
    autoFocus
    className="login"
    iconName="user"
    isOpen
    isCloseButtonShown={false}
    canOutsideClickClose={false}
    canEscapeKeyClose={false}
    title={txt('Authentication')}
  >
    <div className="pt-dialog-body">
      <div className="pt-input-group">
        <label className="pt-label">
          {txt('Username')}
          <input type="text" className="pt-input pt-large" placeholder="eu@email.com" />
        </label>
      </div>
      <div className="pt-input-group">
        <label className="pt-label">
          {txt('Password')}
          <input type="password" className="pt-input pt-large" />
        </label>
      </div>
      <Checkbox onChange={console.log}>{txt('Remember me')}</Checkbox>
    </div>
    <div className="pt-dialog-footer">
      <div className="pt-dialog-footer-actions">
        <a>{txt('Forgot my password')}</a>
        <Button intent={Intent.SUCCESS} text={txt('Login')} />
      </div>
    </div>
  </Dialog>;

export default Login;
