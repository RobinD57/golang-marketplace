<template lang="html">
  <div id="app">
    <header>
      <img src="http://acmelogos.com/images/logo-1.svg" alt="">
      <div class="">
        <Metamask></Metamask>
      </div>
    </header>
  </div>
</template>

<script>

import Metamask from './Metamask'

export default {
  name: 'header-nav',
  props: ['name'],
  components: {
    Metamask
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
  border-bottom: 1px solid #42b983;
  text-align: center;
  background: #ffffff;
  align-items: center;
}
header p {
  padding-right: 4rem;
}

</style>
