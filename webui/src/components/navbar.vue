<template>
  <b-navbar toggleable="md" :sticky="true">

    <b-navbar-brand to="/">
      {{ title }}
    </b-navbar-brand>

    <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

    <b-collapse id="nav-collapse" is-nav>
      <b-navbar-nav>
        <b-nav-item :to="{name:'types'}">
          <b-icon-ui-radios/>
          Types
        </b-nav-item>
        <b-nav-item :to="{name:'content'}">
          <b-icon-file-text-fill/>
          Content
        </b-nav-item>
        <b-nav-item :to="{name:'uploads'}">
          <b-icon-collection-play-fill/>
          Uploaded Media
        </b-nav-item>
        <b-nav-item to="/uploads/" disabled>Disabled</b-nav-item>
      </b-navbar-nav>

      <!-- Right aligned nav items -->
      <b-navbar-nav class="ml-auto">
        <!--
        <b-nav-form>
          <b-form-input size="sm" class="mr-sm-2" placeholder="Search"></b-form-input>
          <b-button size="sm" class="my-2 my-sm-0" type="submit">Search</b-button>
        </b-nav-form>
        -->

        <b-nav-item-dropdown text="Lang" right>
          <b-dropdown-item href="#">EN</b-dropdown-item>
          <b-dropdown-item href="#">ES</b-dropdown-item>
          <b-dropdown-item href="#">RU</b-dropdown-item>
          <b-dropdown-item href="#">FA</b-dropdown-item>
        </b-nav-item-dropdown>

        <b-nav-item-dropdown right>
          <!-- Using 'button-content' slot -->
          <template v-slot:button-content>
            <em>User</em>
          </template>
          <b-dropdown-item href="#">Profile</b-dropdown-item>
          <b-dropdown-item href="#">Sign Out</b-dropdown-item>
        </b-nav-item-dropdown>
      </b-navbar-nav>
    </b-collapse>
  </b-navbar>
</template>

<script>
export default {
  name: "navbar",
  props: {
    title: {
      type: String,
      default: "Gokapi",
    },
  },
}
</script>

<style lang="scss">

</style>
