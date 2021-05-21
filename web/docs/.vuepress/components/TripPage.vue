<template>
  <c-flex direction="column" align="center">
    <c-flex direction="column" h="400px" w="100%">
      <c-flex direction="column" justify="flex-end" flex="1" bg="tomato">
          <c-heading size="md"> 1 Jul — 24 Aug</c-heading>
          <c-heading> Madeira Co-Living</c-heading>
          <c-heading size="md"> Posted by Jack</c-heading>
      </c-flex>
    </c-flex>
    <c-button>
      Chakra Consumed! ⚡️
    </c-button>
  </c-flex>
</template>

<script lang="js">
import {CBox, CButton, CFlex, CHeading, CText} from '@chakra-ui/vue'

export default {
  name: 'TripPage',
  components: {
    CBox,
    CButton,
    CFlex,
    CHeading,
    CText
  }
}
</script>
