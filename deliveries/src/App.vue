<template>
  <div id="app">
    <div class="container" v-for="(item, index) in this.items" :key="index">
      <Delivery :data="item"/>
    </div>
  </div>
</template>

<script>
import Delivery from './components/Delivery.vue'
import {getDataOnFormat} from './services/routes.service'

export default {
  name: 'App',
  components: {
    Delivery
  },
  data() {
    return {
      connection: null,
      items: []
    }
  },
  methods: {
    /**
     * Adquire la información del item que se recibe del WebSocket y la delega a validatePushIsValid para realizar el manejo de la información
     */
    addItem(event) {
      const data = JSON.parse(event.data)
      let obj = getDataOnFormat(data.route_id, data)
      this.validatePushIsValid(obj)
      
    },
    /**
     * Valida que el registro ya este dentro de los items. Si está asigna los valores nuevos al objeto. Si no está lo agrega a items
     */
    validatePushIsValid(object){
      const index = this.items.findIndex(i => i.route_id == object.route_id)
      if(index > -1) {
        let itemChange = this.items[index]
        itemChange.status = object.status
        itemChange.completed_at = object.completed_at
      } else {
        this.items.push(object)
      }

    }
  },
  mounted() {
    this.connection = new WebSocket("ws://localhost:8080")

    this.connection.addEventListener('message', this.addItem);

    this.connection.addEventListener('open', function () {
        console.log('Connection Successfully!');
    });

  }
}
</script>

<style>
  /* http://meyerweb.com/eric/tools/css/reset/ 
   v2.0 | 20110126
   License: none (public domain)
*/
html, body, div, span, applet, object, iframe,
h1, h2, h3, h4, h5, h6, p, blockquote, pre,
a, abbr, acronym, address, big, cite, code,
del, dfn, em, img, ins, kbd, q, s, samp,
small, strike, strong, sub, sup, tt, var,
b, u, i, center,
dl, dt, dd, ol, ul, li,
fieldset, form, label, legend,
table, caption, tbody, tfoot, thead, tr, th, td,
article, aside, canvas, details, embed,
figure, figcaption, footer, header, hgroup,
menu, nav, output, ruby, section, summary,
time, mark, audio, video {
  margin: 0;
  padding: 0;
  border: 0;
  font-size: 100%;
  font: inherit;
  vertical-align: baseline; }

/* HTML5 display-role reset for older browsers */
article, aside, details, figcaption, figure,
footer, header, hgroup, menu, nav, section {
  display: block; }

body {
  line-height: 1; }

ol, ul {
  list-style: none; }

blockquote, q {
  quotes: none; }

blockquote:before, blockquote:after,
q:before, q:after {
  content: '';
  content: none; }

table {
  border-collapse: collapse;
  border-spacing: 0; }

html {
  font-size: 15px;
  font-family: Arial; }

html {
  box-sizing: border-box;
  background-color: #e3e3e3; }

*, *:before, *:after {
  box-sizing: inherit; }

.label {
  font-size: .8rem;
  display: block;
  margin-bottom: .7rem;
  color: #bfbfbf; }

.card {
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  width: auto;
  background-color: white;
  justify-content: center;
  margin: .5rem;
  border-radius: 5px; }
  .card_info {
    margin: 1rem;
    position: relative;
    flex-grow: 1; }
    .card_info--center {
      vertical-align: middle; }
  .card--onroute {
    border-left: 8px solid blue; }
  .card--completed {
    border-left: 8px solid green; }

.button {
  background-color: transparent;
  border: none;
  text-align: center;
  padding: .5rem;
  border-radius: 5px;
  position: absolute;
  top: 50%;
  left: 10%;
  transform: translate(-50%, -50%);
  text-align: center; }
  .button--completed {
    background-color: green;
    color: white;
    width: 6.3rem; }
  .button--onroute {
    background-color: blue;
    color: white;
    width: 70px; }

.container {
  width: 80vw;
  margin: 1rem auto; }

</style>
