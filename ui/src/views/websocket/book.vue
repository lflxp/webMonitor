<template>
  <el-container>
    <el-container style="margin-top:1px;">
      <el-aside width="22%" v-if="showTree">
        <!-- <el-card shadow="always">  -->
          <el-input
            placeholder="输入关键字进行过滤"
            v-model="filterText">
          </el-input>
            <!-- <el-button icon="el-icon-search" circle></el-button> -->
          <div class="custom-tree-container">
            <div class="block">
              <!-- show-checkbox -->
              <el-tree
                ref="tree2"
                :data="data4"
                node-key="id"
                accordion
                draggable
                default-expand-all
                :props="defaultProps"
                :expand-on-click-node="true"
                :filter-node-method="filterNode"
                @node-drag-start="handleDragStart"
                @node-drag-enter="handleDragEnter"
                @node-drag-leave="handleDragLeave"
                @node-drag-over="handleDragOver"
                @node-drag-end="handleDragEnd"
                @node-drop="handleDrop"
                :allow-drop="allowDrop"
                :allow-drag="allowDrag"
                @node-click="handleNodeClick"
                :render-content="renderContent">
              </el-tree>
            </div>
          </div>
        <!-- </el-card> -->
      </el-aside>
        
      <el-main>
        <!--  style="background: pink;" -->
        <div id="app" @contextmenu="showMenu">
          <vue-context-menu :contextMenuData="contextMenuData"
                          @displayitem="displayitem"
                          @savedata="savedata"
                          @showHtml="showHtml">
          </vue-context-menu>
          <!-- <el-card shadow="hover"> -->
            <!-- <el-row style="position: fixed;" > -->
              <!-- <el-button size="mini" icon="el-icon-tickets" @click="showTree = !showTree"></el-button>
              <el-button @click="markdown2Html" type="primary" icon="el-icon-document">To HTML</el-button> -->
            
              <div class="editor-container">
                <markdown-editor id="contentEditor" ref="contentEditor" v-model="content" :height="600" :zIndex="20"></markdown-editor>
              </div>
              <!-- <div v-html="html"></div> -->
              <el-dialog title="预览" :visible.sync="dialogTableVisible">
                <div v-html="html"></div>
              </el-dialog>
            <!-- </el-row> -->
          <!-- </el-card>  -->
        </div>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
  import MarkdownEditor from '@/components/MarkdownEditor'
  import Utils from '@/utils/remote'

  const content = `
  **this is test**
  * vue
  * element
  * webpack
  ## Simplemde
  `
export default {
    name: 'Book',
    components: { MarkdownEditor, Utils },
    data() {
      return {
        dialogTableVisible: false,
        filterText: '',
        tablename: 'tree',
        // data4: JSON.parse(JSON.stringify(data)),
        data4: [],
        current: {},
        defaultProps: {
          children: 'children',
          label: 'label'
        },
        content: content,
        html: '',
        showTree: true,
        // contextmenu data (菜单数据)
        contextMenuData: {
          // the contextmenu name(@1.4.1 updated)
          menuName: 'demo',
          // The coordinates of the display(菜单显示的位置)
          axios: {
            x: null,
            y: null
          },
          // Menu options (菜单选项)
          menulists: [
            {
              fnHandler: 'displayitem', // Binding events(绑定事件)
              icoName: 'fa fa-bars', // icon (icon图标 )
              btnName: '显示/隐藏菜单' // The name of the menu option (菜单名称)
            },
            {
              fnHandler: 'showHtml',
              icoName: 'fa fa-lightbulb-o',
              btnName: '预览'
            },
            {
              fnHandler: 'savedata',
              icoName: 'fa fa-save',
              btnName: '保存'
            }
          ]
        }
      }
    },
    watch: {
      filterText(val) {
        this.$refs.tree2.filter(val)
      }
    },
    mounted() {
      Utils.getIndex().then(data => {
        this.data4 = JSON.parse(data.data)
        console.log('initial', this.data4)
      })
    },
    methods: {
      displayitem() {
        this.showTree = !this.showTree
      },
      showMenu() {
        event.preventDefault()
        var x = event.clientX
        var y = event.clientY
        // Get the current location
        this.contextMenuData.axios = {
          x, y
        }
      },
      showHtml() {
        this.dialogTableVisible = true
        this.markdown2Html()
      },
      savedata() {
        if (this.current.children === undefined) {
          console.log('this is ok')
          // var d2d = { 'tablename': this.tablename, 'key': this.current.label + '_' + this.current.id, 'value': JSON.stringify(this.content) }
          var d2d = { 'tablename': this.tablename, 'key': this.current.label + '_' + this.current.id, 'value': this.content }
          Utils.changeIndex(d2d).then(datas => {
            console.log(datas)
          })
        } else {
          this.$message({
            message: '非叶子节点，不能进行数据保存!',
            type: 'error'
          })
        }
      },
      markdown2Html() {
        import('showdown').then(showdown => {
          const converter = new showdown.Converter()
          this.html = converter.makeHtml(this.content)
        })
      },
      handleNodeClick(data) {
        this.current = data
        if (this.current.children === undefined) {
          console.log('this is ok')
          Utils.checkData(this.tablename, this.current.label + '_' + this.current.id).then(datas => {
            console.log('asdasd', datas)
            // this.content = JSON.stringify(datas.data)
            this.content = datas.data
          })
        }
        console.log(data, new Date().getTime())
      },
      handleDragStart(node, ev) {
        console.log('drag start', node)
      },
      handleDragEnter(draggingNode, dropNode, ev) {
        console.log('tree drag enter: ', dropNode.label)
      },
      handleDragLeave(draggingNode, dropNode, ev) {
        console.log('tree drag leave: ', dropNode.label)
      },
      handleDragOver(draggingNode, dropNode, ev) {
        console.log('tree drag over: ', dropNode.label)
      },
      handleDragEnd(draggingNode, dropNode, dropType, ev) {
        console.log('tree drag end: ', dropNode && dropNode.label, dropType)
      },
      handleDrop(draggingNode, dropNode, dropType, ev) {
        console.log('tree drop: ', dropNode.label, dropType)
      },
      allowDrop(draggingNode, dropNode, type) {
        if (dropNode.data.label === '二级 3-1') {
          return type !== 'inner'
        } else {
          return true
        }
      },
      allowDrag(draggingNode) {
        return draggingNode.data.label.indexOf('三级 3-2-2') === -1
      },
      filterNode(value, data) {
        if (!value) return true
        return data.label.indexOf(value) !== -1
      },
      append(data) {
        var labels = prompt('请输入新增节点名称:')
        const newChild = { id: new Date().getTime(), label: labels }
        if (!data.children) {
          this.$set(data, 'children', [])
        }
        // this.data4 = data.children.push(newChild)
        data.children.push(newChild)

        var d2d = { 'tablename': 'index', 'key': 'index', 'value': JSON.stringify(this.data4) }
        Utils.changeIndex(d2d).then(datas => {
          console.log(datas)
        })
      },

      remove(node, data) {
        const parent = node.parent
        const children = parent.data.children || parent.data
        const index = children.findIndex(d => d.id === data.id)
        children.splice(index, 1)
        // update index
        var d2d = { 'tablename': 'index', 'key': 'index', 'value': JSON.stringify(this.data4) }
        Utils.changeIndex(d2d).then(datas => {
          console.log(datas)
        })
      },

      renderContent(h, { node, data, store }) {
        return (
          <span class='custom-tree-node'>
            <span>{node.label}</span>
            <span class='button'>
              <el-button size='mini' type='text' on-click={ () => this.append(data) }><i class='el-icon-circle-plus'></i></el-button>
              <el-button size='mini' type='text' on-click={ () => this.remove(node, data) }><i class='el-icon-delete'></i></el-button>
            </span>
          </span>)
      }
    }
  }
</script>

<style>
  .custom-tree-node {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 14px;
    padding-right: 8px;
  }

  .custom-tree-node .button {
    display: none;
  }

  .custom-tree-node:hover .button{
    display: block;
  }

  .el-main {
    padding: 0px;
  }

</style>