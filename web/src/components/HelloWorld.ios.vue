<template>
  <GridLayout rows="auto, auto">
    <Label class="message" :text="msg" row="0" horizontalAlignment="center" />
    <!-- copy-webpack-plugin copies asset from src/assets to project output/build directory /assets -->
    <Image src="~/components/icon.png" row="1" class="m-40" />
  </GridLayout>
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
