<template>
  <div class="livehome">
    <!--    <reflv :url="liveAddr" type="flv" islive="" cors="">-->
    <h1 v-if="!is_user">没有这名主播哦.</h1>
    <div v-else>
      <video id="videoElement" controls></video>
    </div>
  </div>
</template>

<script>
  import * as API from '../api'
  import flvjs from 'flv.js'

  export default {
    name: "LiveHome",
    data() {
      return {
        is_user: true,
        liveAddr: "",
        nickname: "",
        playerOptions: {
          fluid: true,
          sources: [{
            type: 'video/flv',
            src: this.liveAddr
          }],
        }
      }
    },

    created() {
      API.getUserInfo(this.$route.params["uid"]).then(response => {
        if (response.data === null) {
          this.is_user = false;
        } else {
          this.liveAddr = response.data.live_addr + ".flv";
          this.nickname = response.data.nickname;
          let flvSrc = this.liveAddr;
          if (flvjs.isSupported()) {
            var videoElement = document.getElementById('videoElement');
            var flvPlayer = flvjs.createPlayer({
              type: 'flv',
              url: flvSrc
            });
            flvPlayer.attachMediaElement(videoElement);
            flvPlayer.load();
            flvPlayer.play()
          }
        }
      });
    },
    mounted() {

    }
  }
</script>

<style scoped>
  video {
    object-fit:fill;
    width:1252px;
    height:669px;
  }
</style>
