<template lang="html">
  <div class="user-container">
    <div class="">
      <p class="user-name">{{user}}</p>
      <div class="user-details">
        <p>User since: 2020</p>
        <p>Listing posted: {{this.postingDate}}</p>
      </div>
    </div>
    <div class="reviews-container">
      <UserReview/>
      <UserReview/>
      <UserReview/>
    </div>
    <button class="msg-button" type="button" name="button">Message</button>
  </div>
</template>

<script>

import UserReview from './UserReview'

export default {
  props: ['user', 'postingDate'],
  name: 'UserSideBar',
  components: {
    UserReview
  },
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
  z-index: 0;
}

.user-name {
  margin-top: 3rem;
  text-align: center;
  text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
  font-weight: bold;
  font-size: 20px;
}

.user-details {
  margin: 2rem 3rem;
}

.user-details p {
  font-size: 12px;
}

.reviews-container {
  border-left: 2px solid rgba(246, 246, 246, 1);
  width: 90%;
  margin: 0 auto;
  border-radius: 5px;
}

.msg-button {
  display: flex;
  margin: 0 auto;
  margin-top: 3rem;
  width: 150px;
  font-family: 'Crimson Text', serif;
  opacity: .7;
  border: none;
  padding: .4rem;
  border-radius: 5px;
  justify-content: center;
  cursor: pointer;
  text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
  box-shadow: 0.5px 0.5px rgba(0, 0, 0, 0.1);
}

.msg-button:hover {
  opacity: .5;
}
</style>
