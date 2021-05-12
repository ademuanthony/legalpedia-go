import { Controller } from 'stimulus'

export default class extends Controller {
  static get targets () {
    return [
      'definitionTitle', 'definitionBody'
    ]
  }

  initialize () {
    this.definitionModal = $('#definitionModal')
  }

  viewDefinition (event) {
    const target = event.currentTarget
    this.definitionTitleTarget.innerText = target.getAttribute('data-title')
    this.definitionBodyTarget.innerHTML = target.getAttribute('data-content')
  }
}
