import { ref, Ref } from 'vue'
const itemBreadcrumbs: Ref = ref([
  {
    title: "Dashboard",
    disabled: false,
    href: "breadcrumbs_dashboard",
  },
  {
    title: "Link 1",
    // disabled: false,
    disabled: true,
    href: "breadcrumbs_link_1",
  },
  // {
  //   title: "Link 2",
  //   disabled: true,
  //   href: "breadcrumbs_link_2",
  // },
])

export {
  itemBreadcrumbs
}
