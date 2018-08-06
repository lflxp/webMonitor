<template>
  <div id="app" @contextmenu="showMenu" style="width: 100px;height: 100px;background: red;">
    <vue-context-menu :contextMenuData="contextMenuData"
	                  @savedata="savedata"
	                  @newdata="newdata">
    </vue-context-menu>
  </div>
</template>
<script>
  export default {
    name: 'app',
    data() {
      return {
        contextMenuData: { // contextmenu data (菜单数据)
        // the contextmenu name(@1.4.1 updated)
          menuName: 'demo',
          axios: { // The coordinates of the display(菜单显示的位置)
            x: null,
            y: null
          },
          menulists: [ // Menu options (菜单选项)
            {
              fnHandler: 'savedata', // Binding events(绑定事件)
              icoName: 'fa fa-home fa-fw', // icon (icon图标 )
              btnName: 'Save' // The name of the menu option (菜单名称)
            },
            {
              fnHandler: 'newdata',
              icoName: 'fa fa-home fa-fw',
              btnName: 'New'
            }
          ]
        }
      }
    },
    methods: {
      showMenu() {
        event.preventDefault()
        var x = event.clientX
        var y = event.clientY
        // Get the current location
        this.contextMenuData.axios = {
          x, y
        }
      },
      savedata() {
        alert(1)
      },
      newdata() {
        console.log('newdata!')
      }
    }
  }
</script>