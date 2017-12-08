<template>
  <div>
    <p>{{ msg }}</p>
  </div>
</template>

<script>
export default {
  name: 'Player',
  data () {
    return {
      ws: null,
      msg: 'hi'
    }
  },
  created () {
    let self = this
    this.ws = new WebSocket('ws://127.0.0.1:8000/ws')
    this.ws.binaryType = 'arraybuffer'
    this.ws.addEventListener('message', function (e) {
      let dataView = new DataView(e.data)
      let decoder = new TextDecoder('utf-8')
      let decodedString = decoder.decode(dataView)
      // let msg = JSON.parse(decodedString)
      self.msg = decodedString

      let encoder = new TextEncoder('utf-8')
      let encodedBuffer = encoder.encode('hi there')
      self.ws.send(encodedBuffer)

    })
  }
}
</script>

<style scoped>

</style>
