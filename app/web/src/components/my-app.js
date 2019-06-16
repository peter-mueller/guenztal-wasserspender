/**
@license
Copyright (c) 2018 The Polymer Project Authors. All rights reserved.
This code may only be used under the BSD style license found at http://polymer.github.io/LICENSE.txt
The complete set of authors may be found at http://polymer.github.io/AUTHORS.txt
The complete set of contributors may be found at http://polymer.github.io/CONTRIBUTORS.txt
Code distributed by Google as part of the polymer project is also
subject to an additional IP rights grant found at http://polymer.github.io/PATENTS.txt
*/
import { LitElement, html, css } from "../../node_modules/lit-element/lit-element.js";
import { setPassiveTouchGestures } from "../../node_modules/@polymer/polymer/lib/utils/settings.js";
import { installRouter } from "../../node_modules/pwa-helpers/router.js";
import { updateMetadata } from "../../node_modules/pwa-helpers/metadata.js"; // These are the elements needed by this element.

import "../../node_modules/@polymer/app-layout/app-header/app-header.js";
import "../../node_modules/@polymer/app-layout/app-toolbar/app-toolbar.js";

class MyApp extends LitElement {
  static get properties() {
    return {
      appTitle: {
        type: String
      },
      _page: {
        type: String
      },
      _drawerOpened: {
        type: Boolean
      },
      _snackbarOpened: {
        type: Boolean
      },
      _offline: {
        type: Boolean
      }
    };
  }

  static get styles() {
    return [css`
        :host {
          display: block;
          user-select: none;

          --app-drawer-width: 256px;

          --app-primary-color: #2757ab;
          --app-secondary-color: #fff699;
          --app-dark-text-color: var(--app-secondary-color);
          --app-light-text-color: white;
          --app-section-even-color: #f7f7f7;
          --app-section-odd-color: white;

          --mdc-theme-primary: var(--app-primary-color);
          --mdc-theme-secondary: var(--app-secondary-color);

          --app-header-background-color: var(--app-primary-color);
          --app-header-text-color: var(--app-light-text-color);
          --app-header-selected-color: var(--app-primary-color);

          --app-drawer-background-color: var(--app-secondary-color);
          --app-drawer-text-color: var(--app-light-text-color);
          --app-drawer-selected-color: #78909C;
        }

        app-header {
          background-color: var(--app-header-background-color);
          color: var(--app-header-text-color);
        }


        .page {
          display: none;
        }

        .page[active] {
          display: block;
        }
      `];
  }

  render() {
    // Anything that's related to rendering should be done in here.
    return html`
      <!-- Header -->
      <app-header>
        <app-toolbar class="toolbar-top">
          <div main-title>${this.appTitle}</div>
        </app-toolbar>
      </app-header>
      
      <!-- Main content -->
      <main role="main" class="main-content">
        <view-control class="page" ?active="${this._page === 'control'}"></view-control>
        <my-view404 class="page" ?active="${this._page === 'view404'}"></my-view404>
      </main>
      
      
    `;
  }

  constructor() {
    super(); // To force all event listeners for gestures to be passive.
    // See https://www.polymer-project.org/3.0/docs/devguide/settings#setting-passive-touch-gestures

    setPassiveTouchGestures(true);
  }

  firstUpdated() {
    installRouter(location => this._locationChanged(location));
  }

  updated(changedProps) {
    if (changedProps.has('_page')) {
      const pageTitle = this.appTitle + ' - ' + this._page;
      updateMetadata({
        title: pageTitle,
        description: pageTitle // This object also takes an image property, that points to an img src.

      });
    }
  }

  _locationChanged(location) {
    const path = window.decodeURIComponent(location.pathname);
    const page = path === '/' ? 'control' : path.slice(1);

    this._loadPage(page); // Any other info you might want to extract from the path (like page type),
    // you can do here.
    // Close the drawer - in case the *path* change came from a link in the drawer.


    this._updateDrawerState(false);
  }

  _updateDrawerState(opened) {
    if (opened !== this._drawerOpened) {
      this._drawerOpened = opened;
    }
  }

  _loadPage(page) {
    switch (page) {
      case 'control':
        import('./view-control.js').then(module => {// Put code in here that you want to run every time when
          // navigating to view1 after my-view1.js is loaded.
        });
        break;

      default:
        page = 'view404';
        import("./my-view404.js");
    }

    this._page = page;
  }

}

window.customElements.define('my-app', MyApp);