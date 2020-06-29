<template lang="html">
  <div id="app">
    <header>
      <img src="http://acmelogos.com/images/logo-1.svg" alt="">
      <div class="">
        <metamask></metamask>
      </div>
    </header>
  </div>
</template>

<script>

import metamask from './Metamask'

export default {
  name: 'header-nav',
  props: ['name'],
  components: {
    metamask
  },
}
</script>



<style lang="css" scoped>

header {
  display: flex;
  justify-content: space-between;
  padding: 0px 1.5rem;
  position: fixed;
  top: 0;
  width: 100%;
  min-height: 90px;
  border-bottom: 2px solid rgba(246, 246, 246, 1);
  text-align: center;
  background: #ffffff;
  align-items: center;
}
header p {
  padding-right: 4rem;
}

</style>
