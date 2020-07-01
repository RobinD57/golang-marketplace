<template lang="html">
  <div class="user-container">
    <div class="">
      <p class="user-name">{{seller.slice(0,3)}}...{{seller.slice(39)}}</p>
      <div class="user-details">
        <p>User since: 2020</p>
        <p>Listing posted: {{this.postingDate}}</p>
      </div>
    </div>
    <div class="reviews-container">
      <UserReview
      :rating='review.rating'
      :content='review.content'
      :reviewer='review.reviewer'
      v-for='review in reviews'
      :key='review.id'/>
    </div>
    <button class="msg-button" type="button" name="button">Message</button>
  </div>
</template>

<script>

import UserReview from './UserReview'

export default {
  props: ['seller', 'postingDate'],
  name: 'UserSideBar',
  components: {
    UserReview
  },
    data() {
    return {
      endpoint: 'http://localhost:8080/user/',
      reviews: []
    }
  },
    methods: {
      async fetchData(address) {
        let res = await fetch(`${this.endpoint}${address}`);
        let data = await res.json();
        console.log(data);
        return this.setResults(data.Reviews);
      },
      setResults(results) {
        this.reviews = results;
      },
    },
    created() {
      this.fetchData(this.seller);
    }
  }
</script>

<style lang="css" scoped>

.user-container{
  width: 350px;
  border: 15px solid rgba(246, 246, 246, .8);
  border-radius: 10px;
  box-shadow: 0.5px 0.5px rgba(0, 0, 0, 0.1);
  z-index: 0;
  margin-bottom: 150px;
}

.user-name {
  margin-top: 3rem;
  text-align: center;
  text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
  font-weight: bold;
  font-size: 20px;
}

.user-details {
  margin: 2rem 1.5rem;
  border-radius: 5px;
}

.user-details p {
  font-size: 14px;
}

.reviews-container {
  border-left: 2px solid rgba(246, 246, 246, 1);
  width: 90%;
  margin: 0 auto;
  border-radius: 5px;
  overflow-y: scroll;
  max-height: 300px;
}

.msg-button {
  display: flex;
  margin: 0 auto;
  margin-top: 3rem;
  margin-bottom: 3rem;
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
