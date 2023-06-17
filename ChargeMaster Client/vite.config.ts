import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server:{
    open :true,
    proxy:{
      '/api':{
        target:"http://121.43.119.64:8848/",
        // target: "http://10.28.216.178:5000/",
        // target: "http://10.128.231.206:8081/",

        changeOrigin:true,
        rewrite(path){
          return path.replace(/^\/api/,'')
        }
      }

    }

  },

  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})
