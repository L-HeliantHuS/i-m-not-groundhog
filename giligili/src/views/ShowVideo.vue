<template>
  <div class="showvideo">
    <!--标题-->
    <el-row :gutter="10">
     <el-col :span="24">
       <span class="title">{{ videoInfo.title }}</span>    ----   <span class="info">{{ videoInfo.info }}</span>
     </el-col>
    </el-row>
    <el-row :gutter="20">
     <el-col :span="12">
       <h5><strong>播放量:</strong><span class="view">{{ videoInfo.view }}</span></h5>
     </el-col>
    </el-row>

    <el-row :gitter="20">
      <el-col :span="17">
        <video-player
          :options="playerOptions">
        </video-player>
      </el-col>
    </el-row>
  </div>

</template>

<script>
  import * as API from '../api'
  import 'video.js/dist/video-js.css'
  import { videoPlayer } from 'vue-video-player'

  export default {
    name: "ShowVideo",
    data() {
      return {
        playerOptions: {
          fluid: true,
          sources: [{
            type: 'video/mp4',
            src: 'http://127.0.0.1/static/lol.mp4'
          }],
        },
        videoInfo: {

        }
      }
    },
    components: {
      videoPlayer
    },
    created() {
      API.getVideo(this.$route.params["videoID"]).then(response => {
        if (response.data === null) {
          this.$router.push('/')
        } else {
          this.videoInfo = response.data;
          this.playerOptions.sources[0].src = response.data.url;
        }
      })
    }
  }
</script>

<style scoped>
  .title {
    font-size: 30px;
  }

  .info {
    font-size: 15px;
  }

  .view {
    color: crimson;
  }
</style>
