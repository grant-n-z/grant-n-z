import {NgModule} from '@angular/core';
import {Routes, RouterModule} from '@angular/router';
import {LoginComponent} from './component/login/login.component';
import {UserIndexComponent} from './component/users/index/user-index.component';
import {LoginRequireGuard} from './guard/login-require-guard.service';
import {LoggingInGuard} from './guard/logging-in-guard.service';
import {GroupIndexComponent} from './component/groups/index/group-index.component';
import {UserPolicyComponent} from './component/users/policy/user-policy.component';
import {GroupUserComponent} from './component/groups/user/group-user.component';
import {GroupPolicyComponent} from './component/groups/policy/group-policy.component';
import {GroupRoleComponent} from './component/groups/role/group-role.component';
import {GroupPermissionComponent} from './component/groups/permission/group-permission.component';
import {ServiceNewComponent} from './component/services/new/service-new.component';


const routes: Routes = [
  {path: '', component: LoginComponent, canActivate: [LoggingInGuard]},

  // Manage service
  {path: 'services/new', component: ServiceNewComponent, canActivate: [LoginRequireGuard]},

  // Per user
  {path: 'users', component: UserIndexComponent, canActivate: [LoginRequireGuard]},
  {path: 'users/policy', component: UserPolicyComponent, canActivate: [LoginRequireGuard]},

  // Per group
  {path: 'groups/:group_id', component: GroupIndexComponent, canActivate: [LoginRequireGuard]},
  {path: 'groups/:group_id/user', component: GroupUserComponent, canActivate: [LoginRequireGuard]},
  {path: 'groups/:group_id/policy', component: GroupPolicyComponent, canActivate: [LoginRequireGuard]},
  {path: 'groups/:group_id/role', component: GroupRoleComponent, canActivate: [LoginRequireGuard]},
  {path: 'groups/:group_id/permission', component: GroupPermissionComponent, canActivate: [LoginRequireGuard]},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
