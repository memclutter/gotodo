import path from 'path'
import { fileURLToPath, URL } from 'url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Inspect from 'vite-plugin-inspect'

import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import Icons from 'unplugin-icons/vite'
import IconsResolver from 'unplugin-icons/resolver'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'

const pathSrc = path.resolve(__dirname, 'src')

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    AutoImport({
      // Auto import functions from Vue, e.g. ref, reactive, toRef...
      imports: ['vue'],

      // Auto import functions from Element Plus, e.g. ElMessage, ElMessageBox... (with style)
      resolvers: [
        ElementPlusResolver(),

        // Auto import icon components
        IconsResolver({
          prefix: 'Icon',
        }),
      ],

      dts: path.resolve(pathSrc, 'auto-imports.d.ts'),
    }),

    Components({
      resolvers: [
        // Auto register icon components
        IconsResolver({
          enabledCollections: ['ep'],
        }),
        // Auto register Element Plus components
        ElementPlusResolver(),
      ],

      dts: path.resolve(pathSrc, 'components.d.ts'),
    }),

    Icons({
      autoInstall: true,
    }),

    Inspect(),
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})
