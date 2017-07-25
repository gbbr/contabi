import React from 'react';
import ReactDOM from 'react-dom';

import { createStore, combineReducers, applyMiddleware } from 'redux'
import { Provider } from 'react-redux'
import createHistory from 'history/createBrowserHistory'
import { Route } from 'react-router'
import { ConnectedRouter, routerReducer, routerMiddleware, push } from 'react-router-redux'

// Create a history of your choosing (we're using a browser history in this case)
const history = createHistory()

// Build the middleware for intercepting and dispatching navigation actions
const middleware = routerMiddleware(history)

// Add the reducer to your store on the `router` key
// Also apply our middleware for navigating
const store = createStore(
  function(state, action) { console.log(action, state); return state; },
  applyMiddleware(middleware)
);

import { App, Invoices, Login } from 'pages'; // eslint-disable-line no-unused-vars

const Dummy = () => <div>Home</div>

document.addEventListener('DOMContentLoaded', () => {
  ReactDOM.render(
    <Provider store={store}>
      <ConnectedRouter history={history}>
        <div>
          <App />
          <Route exact path="/" component={Dummy}/>
          <Route path="/invoices" component={Invoices}/>
          <Route path="/login" component={Login}/>
        </div>
      </ConnectedRouter>
    </Provider>,
    document.getElementById('app-root')
);
});
