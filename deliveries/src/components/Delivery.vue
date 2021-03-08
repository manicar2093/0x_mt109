<template>
  <div class="card"
        :class="cardType"> 
        <div class="card_info">
          <label class="label">Ruta</label>
          <p>{{data.route_id}}</p>
        </div>
        <div class="card_info">
          <label class="label">Conductor</label>
          <p>{{data.driver_name}}</p>
        </div>
        <div class="card_info">
          <label class="label">Creada</label>
          <p>{{data.created_at | formatDate }}</p>
        </div>
        <div class="card_info">
          <label class="label">Inició</label>
          <p>{{data.created_at | formatTime }}</p>
        </div>
        <div class="card_info">
          <label class="label">Terminó</label>
          <p>{{data.completed_at | formatTime }}</p>
        </div>
        <div class="card_info">
          <label class="label">Entregas</label>
          <p>{{data.deliveries}}</p>
        </div>
        <div class="card_info">
          <p class="button"
              :class="buttonType">{{statusText}}</p>
        </div>
      </div>
</template>

<script>
export default {
  name: 'Delivery',
  props: {
    data: Object
  },
  computed: {
    cardType() {
      return this.data.status == 'onroute' ? "card--onroute" : "card--completed"
    },
    buttonType() {
      return this.data.status == 'onroute' ? "button--onroute" : "button--completed"
    },
    statusText() {
      return this.data.status == 'onroute' ? "En Ruta" : "Completada"
    }
  },
  filters: {
    /**
     * Obtiene el formato de tiempo necesario para mostrarlo en el campo
     */
    formatTime(date) {
      if(!date){
        return "-"
      }
      let d = new Date(date)
      let hours = d.getHours()
      let min = d.getMinutes()
      if(hours < 10) {
        hours = `0${hours}`
      }
      if(min < 10) {
        min = `0${min}`
      }
      return `${hours}:${min}`
    },
    /**
     * Obtiene el formato de fecha necesario para mostrarlo en el campo
     */
    formatDate(date) {
      let d = new Date(date)
      let day = d.getDate()
      let month = d.getMonth()
      let year = d.getFullYear()
      const months = [
        'January',
        'February',
        'March',
        'April',
        'May',
        'June',
        'July',
        'August',
        'September',
        'October',
        'November',
        'December'
      ]
      return `${day} de ${months[month]} del ${year}`
    }
  }
}
</script>
