<template lang="html">
  <div class="review-container">
    <div class="left-border"></div>
    <div class="review-details">
      <h3>Good Seller</h3>
      <p>Great prices, and quality pieces, would reccomend</p>
      <p class="rating">4/5</p>
    </div>
  </div>
</template>

<script>

export default {
  props: ['id', 'rating', 'content'],
  name: 'UserReview'
}

</script>

<style lang="css" scoped>

  .review-container {
    display: flex;
    margin: 2rem 0;
    width: 100%;
    align-items: baseline;
    margin: 0 auto;
  }

  .left-border {
    height: 15px;
    border-left: 3px solid black;
    margin-right: 2rem;
  }

  .avatar {
    height: 100%;
    margin: 2rem;
  }

  .user-icon {
    opacity: .5;
    font-size: 20px;
  }

  .review-details {
    font-size: 12px;
  }

  .rating {
    margin-top: .7rem;
  }

  .review-details h3 {
    margin: 0;
  }

  .rating {
    margin-top: .7rem;
  }
</style>
