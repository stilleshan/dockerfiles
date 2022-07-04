<template>
  <div class="box">
    <form method="post" action="#">
      <div class="row gtr-uniform gtr-50">
        <span v-for="(item, index) in inputs" :key="item" class="col-12 row">
          <div
            class="col-10 col-10-mobilep"
            style="text-align: center; padding-top: 10px"
          >
            <input
              type="text"
              v-model.trim="item.inputValue"
              :placeholder="item.placeholder"
            />
          </div>
          <div
            class="col-2 col-2-mobilep"
            style="text-align: center; padding-top: 10px"
          >
            <input
              type="button"
              :value="item.buttonValue"
              :class="item.buttonClass"
              @click="addLinks(index)"
            />
          </div>
        </span>
        <div
          class="col-4 col-4-mobilep list"
          style="text-align: center; padding-top: 20px"
        >
          <select v-model="targetType" @change="selectTarget($event)">
            <option
              v-for="option in targetTypes"
              :key="option"
              :value="option.value"
            >
              {{ option.text }}
            </option>
          </select>
        </div>
        <div
          class="col-4 col-4-mobilep list"
          style="text-align: center; padding-top: 20px"
        >
          <select id="selectApi" v-model="api" @change="selectApi($event)">
            <option v-for="option in apis" :key="option" :value="option.value">
              {{ option.text }}
            </option>
          </select>
        </div>
        <div
          class="col-4 col-4-mobilep"
          style="text-align: center; padding-top: 20px"
        >
          <input
            v-model="manualApiUrl"
            type="text"
            :disabled="isManualApi"
            placeholder="示例：https://sub.ops.ci"
          />
        </div>
        <div
          v-show="isShowMoreConfig"
          class="col-12"
          style="text-align: center; padding-top: 0px"
        >
          <div class="col-12" style="text-align: center; padding-top: 30px">
            <input
              type="text"
              :placeholder="'Include：可选'"
              v-model="moreConfig.include"
            />
          </div>
          <div class="col-12" style="text-align: center; padding-top: 20px">
            <input
              type="text"
              :placeholder="'Exclude：可选'"
              v-model="moreConfig.exclude"
            />
          </div>
          <div class="col-12" style="text-align: center; padding-top: 20px">
            <input
              type="text"
              :placeholder="'远程配置：可选'"
              v-model="moreConfig.remoteconfig"
            />
          </div>
          <div
            class="col-12 col-12-narrower"
            style="text-align: center; padding-top: 20px"
          >
            <input type="checkbox" id="emoji" v-model="moreConfig.emoji" />
            <label for="emoji">Emoji</label>
            <input type="checkbox" id="udp" v-model="moreConfig.udp" />
            <label for="udp">开启UDP</label>
            <input type="checkbox" id="sort" v-model="moreConfig.sort" />
            <label for="sort">排序节点</label>
            <input type="checkbox" id="scv" v-model="moreConfig.scv" />
            <label for="scv">关闭证书检查</label>
            <input type="checkbox" id="nodelist" v-model="moreConfig.list" />
            <label for="nodelist">Node List</label>
          </div>
        </div>
        <div
          class="col-6 col-6-mobilep"
          style="text-align: center; padding-top: 20px"
        >
          <ul class="actions">
            <li>
              <input type="button" value="生成链接" @click="checkAll()" />
            </li>
            <!-- <li><input type="reset" value="重置内容" class="alt" /></li> -->
            <li>
              <input
                type="button"
                value="可选参数"
                class="alt"
                @click="showMoreConfig()"
              />
            </li>
          </ul>
        </div>
        <div class="col-12" style="text-align: center; padding-top: 20px">
          <input
            type="text"
            readOnly="true"
            :placeholder="
              this.apiUrl + '/sub?target=' + this.targetType + '&url='
            "
            v-model.trim="returnUrl"
          />
        </div>
      </div>
    </form>
  </div>
  <DialogLayOut
    :dialogVisible="dialogVisible"
    @closed="resetDialog"
    :message="dialogMessage"
  ></DialogLayOut>
</template>

<script>
import utils from './utils.js';
import DialogLayOut from 'components/common/dialog/DialogLayOut.vue';
export default {
  name: 'HomeForm',
  components: {
    DialogLayOut,
  },
  setup() {
    const ENV = {
      DEFAULT_MORECONFIG: {
        include: '',
        exclude: '',
        remoteconfig: '',
        emoji: true,
        udp: true,
        sort: false,
        scv: false,
        list: false,
      },
    };
    return {
      ENV,
    };
  },
  data() {
    return {
      moreConfig: {},
      isShowMoreConfig: false,
      urls: [],
      returnUrl: '',
      apiUrl: process.env.VUE_APP_BASE_API_URL,
      manualApiUrl: '',
      isManualApi: true,
      api: 'default',
      apis: [
        { value: 'default', text: process.env.VUE_APP_BASE_API_URL },
        { value: 'manual', text: '自定义后端 API 地址' },
      ],
      inputs: [
        {
          buttonClass: '',
          buttonValue: '添加',
          inputValue: '',
          placeholder: '订阅地址',
        },
      ],
      targetType: 'clash',
      targetTypes: [
        { value: 'clash', text: 'Clash' },
        { value: 'clashr', text: 'ClashR' },
        { value: 'v2ray', text: 'V2Ray' },
        { value: 'quan', text: 'Quantumult' },
        { value: 'quanx', text: 'Quantumult X' },
        { value: 'surge&ver=2', text: 'SurgeV2' },
        { value: 'surge&ver=3', text: 'SurgeV3' },
        { value: 'surge&ver=4', text: 'SurgeV4' },
        { value: 'surfboard', text: 'Surfboard' },
        { value: 'ss', text: 'SS (SIP002)' },
        { value: 'sssub', text: 'SS Android' },
        { value: 'ssd', text: 'SSD' },
        { value: 'ssr', text: 'SSR' },
        { value: 'loon', text: 'Loon' },
      ],
      dialogVisible: false,
      dialogMessage: '',
    };
  },
  created() {
    this.moreConfig = this.ENV.DEFAULT_MORECONFIG;
  },
  methods: {
    showMoreConfig() {
      if (this.isShowMoreConfig) {
        this.isShowMoreConfig = false;
      } else {
        this.isShowMoreConfig = true;
      }
    },
    resetDialog() {
      this.dialogVisible = false;
    },
    addLinks(index) {
      if (index == 0) {
        this.inputs.push({
          buttonClass: 'alt',
          buttonValue: '移除',
          placeholder: '订阅地址',
        });
      } else {
        this.inputs.splice(index, 1);
      }
    },
    selectApi(event) {
      if (event.target.value == 'manual') {
        this.isManualApi = false;
      } else {
        this.isManualApi = true;
      }
    },
    selectTarget(event) {
      this.targetType = event.target.value;
    },
    checkUrls() {
      this.urls = [];
      for (const i in this.inputs) {
        if (utils.regexCheck(this.inputs[i].inputValue)) {
          this.urls.push(this.inputs[i].inputValue);
          if (Number(i) + 1 == this.inputs.length) {
            return true;
          }
        } else {
          this.dialogMessage = '请填写正确的订阅地址';
          this.dialogVisible = true;
          return false;
        }
      }
    },
    checkApi() {
      var apiSelect = document.getElementById('selectApi');
      var i = apiSelect.selectedIndex;
      if (apiSelect.options[i].value == 'manual') {
        this.apiUrl = this.manualApiUrl;
        if (!utils.regexCheck(this.apiUrl)) {
          this.dialogMessage = '请填写正确的 API 地址';
          this.dialogVisible = true;
          return false;
        } else if (this.apiUrl.split('').slice(-1) == '/') {
          this.apiUrl = this.apiUrl.substr(0, this.apiUrl.length - 1);
          return true;
        } else {
          return true;
        }
      } else {
        this.apiUrl = apiSelect.options[i].text;
        return true;
      }
    },
    getFinalUrl() {
      this.returnUrl = utils.getSubLink(
        this.urls,
        this.apiUrl,
        this.targetType,
        this.isShowMoreConfig,
        this.moreConfig
      );
    },
    checkAll() {
      if (this.checkUrls() && this.checkApi()) {
        this.getFinalUrl();
      }
    },
  },
};
</script>
