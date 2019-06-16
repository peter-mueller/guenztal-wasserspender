import { html, LitElement, css } from "../../node_modules/lit-element/lit-element.js";

class ObjectTable extends LitElement {
  static get properties() {
    return {
      headers: {
        type: Array
      },
      data: {
        type: Object
      }
    };
  }

  constructor() {
    super();
    this.headers = ['name', 'data.value', 'asdf'];
    this.data = [{
      name: "A",
      data: {
        value: 3
      }
    }, {
      name: "B",
      data: {
        value: 3
      }
    }, {
      name: "A",
      data: {
        value: 3
      }
    }];
  }

  static get styles() {
    return [css`
        div {

                  display: block;
                  overflow-x: auto;
        }
            table {
                    font-family: 'Roboto Slab', serif;
                  border-collapse: collapse;
                  border-top: 1px solid black;
                  border-bottom: 1px solid black;

            }
            td {
                padding: 2px 16px;
            }
            th {
                                padding: 2px 16px;

                border-bottom: 1px solid black;
            }
        `];
  }

  render() {
    return html`
        <div id="scrollcontainer">
            <table>
                <thead>
                    <tr>
                        ${this._headers(this.headers)}
                    </tr>
                </thead >
                <tbody>
                    ${this._rows(this.headers, this.data)}
                </tbody>
            </table >
        </div >
    `;
  }

  _headers(headers) {
    return headers.map(header => {
      return html`<th>${header}</th>`;
    });
  }

  _rows(headers, data) {
    return data.map(datum => {
      return html`<tr>${this._entries(headers, datum)}</tr> `;
    });
  }

  _entries(headers, datum) {
    return headers.map(h => {
      return html`<td>${this._getNestedProperty(datum, h)}</td> `;
    });
  }

  _getNestedProperty(obj, path) {
    return path.split('.').reduce((acc, part) => acc && acc[part], obj);
  }

}

window.customElements.define('object-table', ObjectTable);