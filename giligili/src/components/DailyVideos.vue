<template>
  <div class="dailyvideos">
    <h2 class="top-title">排行榜</h2>
    <el-table :data="videoList" border class="table" :row-class-name="tableRowClassName">
      <el-table-column prop="title" class="title" label="标题" datasrc="url"></el-table-column>
      <el-table-column prop="info" class="info" label="描述" width="100px"></el-table-column>
      <el-table-column prop="view" class="info" label="播放量" width="70px"></el-table-column>
    </el-table>
  </div>
</template>

<script>
  import * as API from '../api'
  export default {
    name: "DailyVideos",
    data() {
      return {
        videoList: []
      }
    },

    created() {
      API.getDailyVideos().then(response => {
        this.videoList = response.data
      })
    },
    methods: {
      tableRowClassName({row, rowIndex}) {
        if (rowIndex === 0) {
          return 'one-row';
        } else if (rowIndex === 1) {
          return 'two-row';
        } else if (rowIndex === 2)  {
          return 'three-row'
        }
        return '';
      }
    }
  }
</script>

<style>
  .top-title{
    display: block;
    text-align: center;
  }
  .table {
    text-overflow: ellipsis;
    white-space: nowrap;
    overflow: hidden;
  }
  .title {
    text-overflow: ellipsis;
    white-space: nowrap;
    overflow: hidden;
  }

  .info {
    text-overflow: ellipsis;
    white-space: nowrap;
    overflow: hidden;
  }

  .el-table .one-row {
    background: #F56C6C;
    color: white;
  }

  .el-table .two-row {
    background: #409EFF;
    color: white;
  }
  .el-table .three-row {
    background: #67C23A;
    color: white;
  }
</style>
