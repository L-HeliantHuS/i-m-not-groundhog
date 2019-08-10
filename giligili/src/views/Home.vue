<template>
  <div class="home">

    <el-row :gutter="20" style="flex: max-content">
      <el-col :xs="24" :sm="12" :md="12" :lg="6" :xl="4" v-for="video in videos" :key="video.id">
        <el-card class="card" @click.native="goVideo(video)">
          <img src="/封面.jpg" class="pic">
          <p class="card-title">
            <span><strong>标题:</strong>{{ video.title }}</span>
          </p>
          <p class="card-info">
            <span><strong>详细:</strong>{{ video.info }}</span>
          </p>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
  import * as API from '../api'

  export default {
    name: 'home',
    components: {},

    data() {
      return {
        videos: []
      }
    },

    methods: {
      goVideo(video) {
        this.$router.push({name: 'showVideo', params: {videoID: video.id}})
      }
    },

    created() {
      this.$notify({
        type: "warning",
        message: "本网站并没有针对移动端进行适配，如果手机访问出现排版混乱属于正常现象。"
      });
      API.getVideos().then((response) => {
        this.videos = response.data;
      })
    }
  }
</script>

<style>
  .pic {
    width: 100%;
  }

  .card {
    margin: 10px;
    position: static;
    top: 70px;
    cursor: pointer;
  }

  .card-title {
    margin-bottom: 10px;
    color: #F56C6C;
    text-overflow: ellipsis;
    white-space: nowrap;
    overflow: hidden;
  }

  .card-info {
    color: #409EFF;
    text-overflow: ellipsis;
    white-space: nowrap;
    overflow: hidden;
  }
</style>
