<script lang="ts" setup>
import { useAuthStore } from "@/stores/auth";
import { User } from "@/types/User";
import { Reactive } from "vue";
import { useDisplay, useTheme } from "vuetify";
const theme = useTheme();
const authStore = useAuthStore();
const drawer = ref(true);
const { smAndDown, md, mdAndUp } = useDisplay();

const initUser: User = {
  id: "",
  username: "",
  password: "",
  full_name: "",
  two_factor_token: "",
  profile_image: "",
  invalid_signIn_count: 0,
  status: false,
  created_at: "",
  updated_at: "",
};

let user: Reactive<User> = reactive(initUser);
const rail: Ref<boolean> = ref(false);
const toggleTheme = () => {
  theme.global.name.value = theme.global.current.value.dark ? "light" : "dark";
};

watchEffect(() => {
  if (smAndDown.value) {
    drawer.value = false;
  }
  if (md.value) {
    rail.value = true;
  }
});

const checkAuth = async () => {
  try {
    if (!authStore.isSignInUserAuth()) {
      // router.replace("/auth/signIn");
      window.location.href = "/auth/signIn";
      return false;
    }

    const userSignIn = await authStore.findUserSignInById(
      authStore.auth.userSignIn.user_id
    );

    Object.assign(user, initUser, userSignIn);
  } catch (error) {
    console.error("checkAuth", error);
    window.location.href = "/auth/signIn";
  }
};

// เช็ก login
checkAuth();
</script>
<template>
  <v-app>
    <!-- display desktop -->
    <v-layout v-if="mdAndUp">
      <v-navigation-drawer
        :location="smAndDown ? 'right' : 'left'"
        v-model="drawer"
        :rail="rail"
        expand-on-hover
        :temporary="smAndDown"
        :permanent="mdAndUp"
        @click="rail = false"
      >
        <template v-slot:prepend>
          <v-list-item
            class="border-b py-3 cursor-pointer"
            lines="two"
            prepend-avatar="./../assets/user.png"
            append-icon="mdi-menu-down"
            :title="user.full_name"
          >
          </v-list-item>
        </template>

        <SidebarComponent />

        <template v-slot:append>
          <v-list density="compact" nav class="border-t">
            <v-list-item
              prepend-icon="mdi-logout"
              title="ออกจากระบบ"
              @click="authStore.handleSignOut()"
              value="logout"
            ></v-list-item>
          </v-list>
        </template>
      </v-navigation-drawer>
      <v-app-bar :elevation="2">
        <template v-slot:prepend>
          <v-app-bar-nav-icon @click="rail = !rail"></v-app-bar-nav-icon>
        </template>
        <v-app-bar-title
          >ระบบการจัดการรายการงาน (Task Management System)</v-app-bar-title
        >
        <template v-slot:append>
          <v-btn
            v-if="theme.global.name.value === 'dark'"
            icon="mdi-white-balance-sunny"
            @click="toggleTheme"
            color=""
          ></v-btn>
          <v-btn
            v-else
            icon="mdi-moon-waning-crescent"
            @click="toggleTheme"
            color=""
          ></v-btn>
        </template>
      </v-app-bar>

      <!-- display content -->
      <v-main class="w-full h-full bg-surface-variant">
        <router-view />
      </v-main>
      <!-- display content -->
    </v-layout>
    <!-- display desktop -->

    <!-- display mobile -->
    <v-layout v-if="smAndDown" class="w-screen h-screen">
      <v-navigation-drawer
        location="right"
        v-model="drawer"
        :rail="rail"
        expand-on-hover
        :temporary="smAndDown"
        :permanent="mdAndUp"
        @click="rail = false"
      >
        <template v-slot:prepend>
          <v-list-item
            append-icon="mdi-menu-down"
            class="border-b py-3 cursor-pointer"
            lines="two"
            prepend-avatar="./../assets/user.png"
            subtitle="Logged in"
            title="Name User"
          ></v-list-item>
        </template>

        <SidebarComponent />

        <template v-slot:append>
          <v-list density="compact" nav class="border-t">
            <v-list-item
              prepend-icon="mdi-logout"
              title="ออกจากระบบ"
              @click="authStore.handleSignOut()"
              value="logout"
            ></v-list-item>
          </v-list>
        </template>
      </v-navigation-drawer>
      <v-main class="w-full h-full bg-surface-variant overflow-y-scroll">
        <div>
          <router-view />
          <v-bottom-navigation grow class="w-full h-full flex-0">
            <v-btn value="tasks" to="/tasks">
              <v-icon>mdi-view-dashboard</v-icon>
              <span class="mt-0.5">จัดการงาน</span>
            </v-btn>
            <v-btn value="profile" to="/profile">
              <v-icon>mdi-account</v-icon>
              <span class="mt-0.5">ข้อมูลผู้ใช้</span>
            </v-btn>
            <v-btn value="drawer" @click="drawer = !drawer">
              <v-icon>mdi-menu</v-icon>
              <span class="mt-0.5">เมนู</span>
            </v-btn>
          </v-bottom-navigation>
        </div>
      </v-main>
    </v-layout>
    <!-- display mobile -->
  </v-app>
</template>
