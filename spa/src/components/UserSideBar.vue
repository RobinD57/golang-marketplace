<template lang="html">
  <div class="user-container">
    <div class="">
      <p class="user-name">{{user}}</p>
    </div>

  </div>
</template>

<script>
export default {
  props: ['user'],
  name: 'UserSideBar',
  data() {
    return {
    }
  }
}

</script>

<style lang="css" scoped>

.user-container{
  width: 40%;
  border: 2px solid rgba(246, 246, 246, 1);
  border-radius: 5px;
}

.user-name {

}
</style>
