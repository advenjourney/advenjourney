<template>
  <div>
    <div class="message">{{ msg }}</div>
    <img src="~/components/icon.png" alt="logo" />
  </div>
</template>
<script>
export default {
  name: "HelloWorld",
  props: {
    msg: String
  }
};
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
