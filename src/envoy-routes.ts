import { html, css, LitElement } from "lit"
import { customElement, property } from "lit/decorators.js"
import "@lrnwebcomponents/lrn-table"

const isDevMode = import.meta.env.DEV

/**
 * An example element.
 *
 * @slot - This element has a slot
 * @csspart button - The button
 */
@customElement("envoy-routes")
export class EnvoyRoutes extends LitElement {

  static styles = css`
    :host {
      display: block;
      border: solid 1px gray;
      padding: 16px;
      max-width: 800px;
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
    const rows = [];
    for (const v of this.envoyData.configs) {
      rows.push(
        html`
          <ui5-table-row type="Active">
            <ui5-table-cell>${v["@type"]}</ui5-table-cell>
            <ui5-table-cell>${v.last_updated}</ui5-table-cell>
          </ui5-table-row>
        </ui5-table>
        `
      )
    }

    return html`
      <span>${this.message}</span>

      <ui5-table >
      <ui5-table-column slot="columns">
        <span>@type</span>
      </ui5-table-column>
      <ui5-table-column slot="columns">
        <span>last_updated</span>
      </ui5-table-column>
      ${rows}
      `
  }
}

// 不要?
// declare global {
//   interface HTMLElementTagNameMap {
//     "envoy-routes": EnvoyRoutes
//   }
// }
