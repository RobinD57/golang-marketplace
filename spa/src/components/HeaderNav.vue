<template lang="html">
  <div id="app">
    <header>
      <div class="header-contents">
        <img src="http://acmelogos.com/images/logo-1.svg" alt="">
        <div class="main-right">
          <div class="right-links">
            <button class='newlisting-button' type="button" name="button">post ad</button>
          </div>
          <div class="metamask">
            <Metamask class="alignment"></Metamask>
          </div>
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

.main-right {
  display: flex;
  align-items: center;
}

.right-links {
  position: relative;
}

.newlisting-button {
  display: flex;
  align-items: center;
  width: 75px;
  height: 25px;
  font-family: 'Crimson Text', serif;
  font-size: 12px;
  font-weight: bold;
  opacity: .7;
  border: none;
  border-radius: 5px;
  justify-content: center;
  cursor: pointer;
  text-shadow: 1px 1px 1px rgba(0,0,0,0.2);
  box-shadow: 0.5px 0.5px rgba(0, 0, 0, 0.1);
  outline: none;
  margin-right: 2rem;
}

.newlisting-button:hover {
  opacity: .5;
}

header p {
  padding-right: 4rem;
}

</style>
