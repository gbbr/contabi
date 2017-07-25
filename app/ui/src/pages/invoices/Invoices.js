import './invoices.scss';

import React, { Component } from 'react';
import { Intent, Button } from '@blueprintjs/core';
import txt from 'lang';

export default class Invoices extends Component {
  componentDidMount() {
    console.log('mounted');
  }

  componentWillUnmount() {
    console.log('unmounted');
  }

  render() {
    return (
      <section className="page page-invoices">
        <Button intent={Intent.SUCCESS} className="btn-add-invoice" iconName="add">{txt('Add invoice')}</Button>
        <table className="invoice-list pt-table pt-striped pt-interactive">
          <thead>
            <tr>
              <th>{txt('Date')}</th>
              <th>{txt('Due by')}</th>
              <th>{txt('Series')}</th>
              <th>{txt('Client')}</th>
              <th>{txt('Amount')}</th>
              <th>{txt('Paid')}</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>21/07/2016</td>
              <td>21/08/2016</td>
              <td>SF-03</td>
              <td>Serfitrans SRL</td>
              <td>300 RON</td>
              <td><span className="color-warning">NU</span></td>
            </tr>
            <tr>
              <td>20/06/2016</td>
              <td>20/07/2016</td>
              <td>SF-02</td>
              <td>Serfitrans SRL</td>
              <td>320 RON</td>
              <td>320 RON</td>
            </tr>
            <tr>
              <td>02/04/2016</td>
              <td>02/05/2016</td>
              <td>SF-01</td>
              <td>Serfitrans SRL</td>
              <td>500 RON</td>
              <td>500 RON</td>
            </tr>
          </tbody>
        </table>
      </section>
    );
  }
}
