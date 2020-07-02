<template lang="html">
  <div class="aside-card-wrapper" v-bind:data-id='this.id'>
    <div class="main-card" v-bind:style= "{ backgroundImage: `url(${this.photo}` }">
    </div>
    <div class="blur">
      <div class="card-details">
        <span>
          {{ this.name.length > 15 ? this.name.slice(0,15) + "..." : this.name }}
        </span>
        <span>${{this.price}}</span>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ['id', 'name', 'price', 'photo'],
}
</script>

<style lang="css" scoped>
  .main-card {
    background-size: cover;
    height: 200px;
    width: 70%;
    margin: 0 auto;
    margin-top: 30px;
    border-radius: 5px;
    box-shadow: 0.5px 0.5px rgba(0, 0, 0, 0.1);
  }

  .blur {
    background: rgba(245, 245, 245, 0.9);
    margin: 0 auto;
    height: 50px;
    width: 70%;
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
