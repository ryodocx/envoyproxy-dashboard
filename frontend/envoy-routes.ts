import { html, css, LitElement } from "lit"
import { customElement, property } from "lit/decorators.js"

import '@vaadin/grid/theme/material/vaadin-grid.js'
import '@vaadin/grid/theme/material/vaadin-grid-filter-column.js'
// import '@vaadin/grid/theme/material/vaadin-grid-selection-column.js'
import '@vaadin/grid/theme/material/vaadin-grid-sort-column.js'
// import '@vaadin/grid/theme/material/vaadin-grid-tree-column.js'

const isDevMode = import.meta.env.DEV

@customElement("envoy-routes")
export class EnvoyRoutes extends LitElement {

  static styles = css`
    :host {
      display: block;
      border: solid 1px gray;
      padding: 16px;
    }
  `

  @property({ type: String })
  message = ""

  @property()
  envoyData = {}

  constructor() {
    super()

    fetch(".vscode/envoy.json")
      .then(response => {
        isDevMode && console.log("envoy.json response", response)
        this.message = `fetch envoy.json -> ${response.statusText}`
        return response.json()
      })
      .then(data => {
        isDevMode && console.log("envoy.json data", data)
        this.envoyData = data
      })
  }

  render() {
    return html`
      <vaadin-grid .items="${this.envoyData.configs[1].dynamic_active_clusters}" theme="column-borders" all-rows-visible>
        <vaadin-grid-filter-column path="cluster.name" auto-width ></vaadin-grid-filter-column>
        <vaadin-grid-filter-column path="last_updated" auto-width></vaadin-grid-filter-column>
        <vaadin-grid-filter-column path="version_info" auto-width></vaadin-grid-filter-column>
      </vaadin-grid>
      `
  }
}

// 不要?
// declare global {
//   interface HTMLElementTagNameMap {
//     "envoy-routes": EnvoyRoutes
//   }
// }
