<template lang="html">
  <div class="user-container">
    <div class="">
      <p class="user-name">{{user}}</p>
      <div class="user-details">
        <p>User since: 2020</p>
        <p>Listing posted: {{posting-date}}</p>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  props: ['user', 'posting-date'],
  name: 'UserSideBar',
  data() {
    return {
    }
  }
}

</script>

<style lang="css" scoped>

.user-container{
  width: 30%;
  border: 2px solid rgba(246, 246, 246, 1);
  border-radius: 10px;
  box-shadow: 0.5px 0.5px rgba(0, 0, 0, 0.1);
}

.user-name {
  margin-top: 3rem;
  text-align: center;
  text-shadow: 1px 1px 3px rgba(0,0,0,0.2);
  font-weight: bold;
  font-size: 20px;
}

.user-details {
  margin: 2rem 3rem;
}

.user-details p {
  font-size: 12px;
}
</style>
