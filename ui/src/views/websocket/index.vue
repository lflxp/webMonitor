<template>
    <el-table
    :data="data"
    style="width: 100%">
    <el-table-column type="expand">
      <template slot-scope="props">
        <el-form label-position="left" inline class="demo-table-expand">
          <el-form-item label="源MAC">
            <span>{{ props.row.srcmac }}</span>
          </el-form-item>
          <el-form-item label="目的MAC">
            <span>{{ props.row.dstmac }}</span>
          </el-form-item>
          <el-form-item label="Headers">
            <li v-for="(value,key) in props.row.headers" :key="value">
              <el-tag type="success">{{ key }}</el-tag>  {{ value }}
            </li>
          </el-form-item>
          <el-form-item label="Response">
            <span>{{ props.row.response }}</span>
          </el-form-item>
        </el-form>
      </template>
    </el-table-column>
    <el-table-column
      label="时间"
      prop="time">
    </el-table-column>
    <el-table-column
      label="Protocol"
      sortable
      prop="protocol">
    </el-table-column>
    <el-table-column
      label="类型"
      prop="type">
    </el-table-column>
    <el-table-column
      label="Method"
      >
      <template slot-scope="scope">
        <span v-if="scope.row.type === 'response'">{{ scope.row.statusCode }} {{ scope.row.statusMsg}}</span>
        <span v-else>{{ scope.row.method }}</span>
      </template>
    </el-table-column>
    <el-table-column
      label="URL"
      prop="url">
    </el-table-column>
    <el-table-column
      label="VERSION"
      prop="version">
    </el-table-column>
    <el-table-column
      label="源IP"
      prop="srcip">
    </el-table-column>
    <el-table-column
      label="源端口"
      prop="srcport">
    </el-table-column>
    <el-table-column
      label="目的IP"
      prop="Dstip">
    </el-table-column>
    <el-table-column
      label="目的端口"
      prop="dstport">
    </el-table-column>
  </el-table>
</template>

<script>
  export default {
    data() {
      return {
        websock: null,
        over: null,
        data: []
      }
    },
    created() {
      // 页面刚进入时开启长连接
      this.initWebSocket()
    },
    destroyed: function() {
      // 页面销毁时关闭长连接
      this.websocketclose()
    },
    methods: {
      initWebSocket() { // 初始化weosocket
        // const wsuri = process.env.WS_API + '/websocket/threadsocket' // ws地址
        const wsuri = 'ws://localhost:8881/ws'
        this.websock = new WebSocket(wsuri)
        this.websock.onopen = this.websocketonopen
        this.websock.onerror = this.websocketonerror
        this.websock.onmessage = this.websocketonmessage
        this.websock.onclose = this.websocketclose
        this.over = () => {
          this.websock.close()
        }
      },
      websocketonopen() {
        console.log('WebSocket连接成功')
      },
      websocketonerror(e) { // 错误
        console.log('WebSocket连接发生错误')
      },
      websocketonmessage(e) { // 数据接收
        console.log(e)
        // const redata = JSON.parse(e.data)
        // 注意：长连接我们是后台直接1秒推送一条数据，
        // 但是点击某个列表时，会发送给后台一个标识，后台根据此标识返回相对应的数据，
        // 这个时候数据就只能从一个出口出，所以让后台加了一个键，例如键为1时，是每隔1秒推送的数据，为2时是发送标识后再推送的数据，以作区分
        // console.log(redata.value)
        if (this.data.length > 200) {
          this.data = []
        }
        this.data.push(JSON.parse(e.data))
      },
      websocketsend(agentData) { // 数据发送
        this.websock.send(agentData)
      },
      websocketclose(e) { // 关闭
        // console.log('connection closed (' + e.code + ')')
        this.initWebSocket()
        console.log('connection close ', e)
      }
    },
    beforeDestroy() {
      this.over()
    }
}
 </script>

 <style>
  .demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
</style>
