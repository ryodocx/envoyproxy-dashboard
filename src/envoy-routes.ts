import { html, css, LitElement } from "lit"
import { customElement, property } from "lit/decorators.js"
// import "@ui5/webcomponents/dist/Table.js"
// import "@ui5/webcomponents/dist/TableColumn.js"
// import "@ui5/webcomponents/dist/TableRow.js"
// import "@ui5/webcomponents/dist/TableCell.js"

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

  /**
   * The name to say "Hello" to.
   */
  @property()
  name = "World"

  /**
   * The number of times the button has been clicked.
   */
  @property({ type: Number })
  count = 0

  @property({ type: Map })
  envoyData = {}

  @property({ type: Number })
  envoyDataCount = 0

  constructor() {
    super();
    this.count = 3

    fetch(".vscode/envoy.json")
      .then(response => {
        import.meta.env.DEV && console.log("envoy.json response", response);
      })
      .then(data => {
        // handle success
        console.log("envoy.json data", data)
        // let s = Object.keys(data)
        this.envoyDataCount = 5
      })
  }

  render() {
    return html`
      <h1>Hello, ${this.name}!</h1>
      <button @click=${this._onClick} part="button">
        Click Count: ${this.count}
      </button>
      data: ${this.envoyDataCount}
      <slot></slot>
    `
  }


  private _onClick() {
    this.count++
  }

  foo(): string {
    return "foo"
  }
}

declare global {
  interface HTMLElementTagNameMap {
    "envoy-routes": EnvoyRoutes
  }
}
