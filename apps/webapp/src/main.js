import './assets/app.scss'
import { mount } from 'svelte'
import App from './App.svelte'

console.log("window.location", window.location)
const app = mount(App, {
  target: document.getElementById('app'),
  props: {
    url: window.location.pathname
  }
})

export default app
