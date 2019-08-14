import { html, LitElement, css } from "../../node_modules/lit-element/lit-element.js";
import "../../node_modules/@material/mwc-button/mwc-button.js";
import "../../node_modules/@material/mwc-icon/mwc-icon.js";
import { sharedStyles } from "./shared-styles.js";

class WaterToggle extends LitElement {
  static get properties() {
    return {
      type: {
        type: String
      },
      label: {
        type: String
      },
      locked: {
        type: Boolean
      },
      active: {
        type: Boolean
      },
      imageSrc: {
        type: String
      }
    };
  }

  static get styles() {
    return [sharedStyles, css`
            .card {
                background-color: white;   
                border-radius: 4px;  
                
                display: flex;
                flex-direction: row;
                align-items: center;


                box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
            }

            .card img {
              display: flex;
                object-fit: cover;
                width: 178px;
                height: 64px;
                transition: filter 1s;
            }
            .card img[locked] {
                filter: grayscale(71%);
            }
            .card section[label] {
                padding: 16px;
                                flex-grow: 1;

            }

            .card section[actions] {
                padding: 8px;
            }

            #label {
                font-size: 1em;
            }

        `];
  }

  constructor() {
    super();
    this.label = "Wasser";
    this.active = false;
  }

  render() {
    return html`
        <div class="card">

        <div>
            <img src=${this.imageSrc} ?locked=${this.locked}>
        </div>
            <section label>
    <div id="label" class="mdc-typography--overline">${this.label}</div>
            </section>

            <section actions>
                <mwc-button
                    ?outlined=${!this.active}
                    ?raised=${this.active}
                    ?disabled=${this.locked}

                    icon=${this.locked ? "lock" : ""}
                    label="${this.active ? "ausschalten" : "einschalten"}"

                    @click=${this._onClick}
                ></mwc-button>

            </section>

        </div>

        `;
  }

  _onClick() {
    let event = new CustomEvent('toggle', {
      detail: {
        type: this.type,
        detail: {
          composed: true,
          bubbles: true
        }
      }
    });
    this.dispatchEvent(event);
  }

}

window.customElements.define('water-toggle', WaterToggle);