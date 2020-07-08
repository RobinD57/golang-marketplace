<template lang="html">
  <div class="aside-card-wrapper" v-bind:data-id='this.id'>
    <div class="main-card" v-bind:style= "{ backgroundImage: `url(${this.photos[0]}` }">
    </div>
    <div class="blur">
      <div class="card-details">
        <span>
          {{ this.name.length > 25 ? this.name.slice(0,25) + "..." : this.name }}
        </span>
        <span>${{this.price}}</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    id: String,
    name: String,
    price: Number,
    photos: Array
  }
}
</script>

<style lang='css' scoped>

  .aside-card-wrapper {
    width: 100%;
    justify-content:
  }

  .main-card {
    display: flex;
    background-size: cover;
    height: 200px;
    width: 80%;
    margin-top: 30px;
    border-radius: 5px;
    box-shadow: 0.5px 0.5px rgba(0, 0, 0, 0.1);
  }

  .blur {
    background: rgba(245, 245, 245, 0.9);
    height: 50px;
    width: 80%;
    margin-top: -4px;
    padding: 2px;
    border-radius: 5px;
  }

  .card-details {
    position: relative;
    display: flex;
    justify-content: space-around;
    margin-top: 2.5px;
  }

  .card-details span {
    text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
    margin-top: 0.8rem;
  }

</style>
