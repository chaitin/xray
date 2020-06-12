<template>
  <div>

    <a-card style="width: 100%"
            :body-style="{padding: 0}">
      <h3 slot="title">Subdomains</h3>
      <div slot="extra" style="width: 600px">
        <a-switch
          style="margin-right: 12px;"
          @change="onSwitchChange"
          checked-children="Web"
          un-checked-children="All"/>
        <a-input-search style="width: 370px; margin-right: 30px"
                        ref="searchWord"
                        rel="abc"
                        @change="onSearchChange"
                        @search="onSearchChange"
                        placeholder="Search Target"></a-input-search>
        <a-button type="primary" icon="reload" @click="window.location.reload()">Reload</a-button>
      </div>
      <a-table :dataSource="subdomainData"
               :expandRowByClick="true"
               rowKey="id"
               :loading="tableLoading"
               @change="onTableChange"
               :pagination="false">
        <div slot="expandedRowRender"
             slot-scope="record"
             style="margin: 0">

          <a-descriptions bordered
                          class="expand-detail"
                          :column="1"
                          size="middle">
            <a-descriptions-item label="Website"
                                 v-if="record.web && record.web.length"
                                 :span="1">
              <template v-for="item in record.web">
                <div :key="item.link">
                  <pre><a :href="item.link" target="_blank">{{item.link}}</a>
Title:   {{item.title}}
Status:  {{item.status}}
Server:  {{item.server}}
                  </pre>
                </div>
              </template>
            </a-descriptions-item>
            <a-descriptions-item label="CNAME" :span="1" v-if="record.cname && record.cname.length">
              <pre>{{record.cname}}</pre>
            </a-descriptions-item>
            <a-descriptions-item label="IPInfo" :span="1" v-if="record.ip && record.ip.length">
              <pre>{{record.ip}}</pre>
            </a-descriptions-item>
            <a-descriptions-item label="Extra" :span="1" v-if="record.extra">
              <pre>{{record.extra}}</pre>
            </a-descriptions-item>
          </a-descriptions>
        </div>
        <a-table-column
          title="ID"
          width="64px"
          :customRender="(text, record, index)=> index+1"
          key="id">
        </a-table-column>
        <a-table-column
          title="Domain"
          dataIndex="domain"
          :sorter="(a,b)=>a.domain.localeCompare(b.domain)">
        </a-table-column>
        <a-table-column
          title="Title"
          key="title">
          <template slot-scope="record">
            {{record.web ? (record.web.length > 0 ? record.web[0].title:""):""}}
          </template>
        </a-table-column>
        <a-table-column
          title="Status"
          class="filter-column"
          :filters="statusTypes"
          @filter="onFilterStatus"
          key="status">
          <template slot-scope="record">
            {{record.web ? (record.web.length > 0 ? record.web[0].status:""):""}}
          </template>
        </a-table-column>
        <a-table-column
          title="Server"
          key="server">
          <template slot-scope="record">
            {{record.web ? (record.web.length > 0 ? record.web[0].server:""):""}}
          </template>
        </a-table-column>
        <a-table-column
          title="Source"
          class="filter-column"
          :filters="sourceTypes"
          @filter="(value, record) => record.verbose_name.includes(value)"
          dataIndex="verbose_name">
        </a-table-column>
      </a-table>
    </a-card>
  </div>
</template>

<script>
  export default {
    name: "subdomain",
    mounted () {
      this.init(false)
    },
    props: {
      data: Array,
      loading: Boolean,
    },
    data () {
      return {
        window: window,
        subdomainData: [],
        tableLoading: this.loading,
      }
    },
    methods: {
      init (webonly) {
        this.tableLoading = true
        let domainData= []
        for (let d of this.data) {
          if (webonly && d.web.length === 0) {
              continue
          }
          let websiteTitle = '';
          for (let w of d.web) {
            websiteTitle += w.title + ' ' + w.server + ' '
          }
          d.searchCache = d.domain + ' ' + d.cname.join(' ') + ' ' + websiteTitle
          domainData.push(d)
        }
        setTimeout(() => {
          this.subdomainData = domainData
          this.tableLoading = false
        }, 500)
      },
      onSearchChange () {
        // 直接用 v-model 在大数据量下太卡了, hack 一下
        let toSearch = this.$refs.searchWord.$el.firstChild.value
        let self = this
        this.tableLoading = true
        let fn = function() {
          let newdata = []
          for (let d of self.data) {
            if (d.searchCache.includes(toSearch)) {
              newdata.push(d)
            }
          }
          self.subdomainData = newdata
          self.$nextTick(() => {
            setTimeout(() => {
              self.tableLoading = false
            }, 100)
          })
        }
        this.$debounce(fn, 800)()
      },
      onFilterStatus (value, record) {
        for (let w of record.web) {
          if (w.status.toString() === value) {
            return true
          }
        }
        return false
      },
      onSwitchChange(checked) {
        return this.init(checked)
      },
      onTableChange() {
        this.tableLoading = true
        this.$nextTick(() => {
          setTimeout(() => {
            this.tableLoading = false
          }, 500)
        })
      }
    },
    computed: {
      statusTypes () {
        let s = new Set();
        for (let d of this.subdomainData) {
          for (let w of d.web) {
            // 直接用 status 会很卡，原因不明
            s.add(w.status.toString())
          }
        }
        let result = []
        for (let t of s.values()) {
          result.push({text: t, value: t})
        }
        return result
      },
      sourceTypes () {
        let s = new Set();
        for (let d of this.subdomainData) {
          s.add(d.verbose_name)
        }
        let result = []
        for (let t of s.values()) {
          result.push({text: t, value: t})
        }
        return result
      },
    },
    watch: {
      data (newData, old) {
        this.init(false)
      },
      loading (newData) {
        this.tableLoading = newData
      }
    }
  }
</script>