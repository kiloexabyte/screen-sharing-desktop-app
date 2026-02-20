import { createRouter, createWebHashHistory } from "vue-router";
import Home from "../pages/index.vue";
import Room from "../pages/room.vue";

const routes = [
  { path: "/", name: "Home", component: Home },
  { path: "/room", name: "Room", component: Room },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
