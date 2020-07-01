<template lang="html">
  <div id="app">
    <header>
      <div class="header-contents">
        <img src="http://acmelogos.com/images/logo-1.svg" alt="">
        <div class="">
          <Metamask class="alignment"></Metamask>
        </div>
      </div>
    </header>
  </div>
</template>

<script>

import Metamask from './Metamask';

export default {
  name: 'header-nav',
  components: {
    Metamask
  },
}
</script>



<style lang="css" scoped>

header {
  padding: 0px 1.5rem;
  display: flex;
  position: fixed;
  top: 0;
  width: 100%;
  min-height: 90px;
  border-bottom: 2px solid rgba(246, 246, 246, 1);
  text-align: center;
  background: #ffffff;
  z-index: 1;
  align-items: center;
}

.header-contents {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

header p {
  padding-right: 4rem;
}

.alignment {
  margin-right: 30%;
}

</style>
