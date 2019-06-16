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

                overflow: hidden;
                width: 256px;

                box-shadow: 0 1px 3px rgba(0,0,0,0.12), 0 1px 2px rgba(0,0,0,0.24);
            }

            .card img {
                object-fit: cover;
                height: 194px;
                width: 100%;
                transition: filter 1s;
            }
            .card img[locked] {
                filter: grayscale(71%);
            }
            .card section {
                padding: 16px;
            }

            .card section[actions] {
                padding: 8px;
                display: flex;
                flex-direction: row-reverse;
                align-items: center;
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
            <img src=${this.imageSrc} ?locked=${this.locked}>

            <section>
    <div class="mdc-typography--overline">${this.label}</div>
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