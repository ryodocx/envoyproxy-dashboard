import { html, css, LitElement } from "lit"
import { customElement, property } from "lit/decorators.js"
import "@ui5/webcomponents/dist/Table.js"
import "@ui5/webcomponents/dist/TableColumn.js"
import "@ui5/webcomponents/dist/TableRow.js"
import "@ui5/webcomponents/dist/TableCell.js"

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

  @property({ type: Map })
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
    let rows = `
      <ui5-table-row>
        <ui5-table-cell>Notebook Basic 15HT-1000</ui5-table-cell>
        <ui5-table-cell>Very Best Screens</ui5-table-cell>
        <ui5-table-cell>30 x 18 x 3cm</ui5-table-cell>
        <ui5-table-cell>4.2KG</ui5-table-cell>
        <ui5-table-cell>956EUR</ui5-table-cell>
      </ui5-table-row>
    </ui5-table>
    `

    return html`
      <span>${this.message}</span>

      <ui5-table>
      <ui5-table-column slot="columns">
        <span>Item</span>
      </ui5-table-column>
      <ui5-table-column slot="columns">
        <span>Item2</span>
      </ui5-table-column>
      ${rows}
      `
  }
}

declare global {
  interface HTMLElementTagNameMap {
    "envoy-routes": EnvoyRoutes
  }
}
