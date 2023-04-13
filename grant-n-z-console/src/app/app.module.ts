import {BrowserModule} from '@angular/platform-browser';
import {CUSTOM_ELEMENTS_SCHEMA, NgModule} from '@angular/core';
import {MatToolbarModule} from '@angular/material/toolbar';
import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {MatIconModule} from '@angular/material/icon';
import {MatSidenavModule} from '@angular/material/sidenav';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatListModule} from '@angular/material/list';
import {LoginComponent} from './component/login/login.component';
import {UserIndexComponent} from './component/users/index/user-index.component';
import {MatCardModule} from '@angular/material/card';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatInputModule} from '@angular/material/input';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {MatButtonModule} from '@angular/material/button';
import {HttpClientModule} from '@angular/common/http';
import {CookieService} from 'ngx-cookie-service';
import {UserService} from './service/user.service';
import {OverlayModule} from '@angular/cdk/overlay';
import {MatSpinner} from '@angular/material/progress-spinner';
import {ToastrModule} from 'ngx-toastr';
import {LoginRequireGuard} from './guard/login-require-guard.service';
import {LoggingInGuard} from './guard/logging-in-guard.service';
import {ServiceService} from './service/service.service';
import {MatTableModule} from '@angular/material/table';
import {MatMenuModule} from '@angular/material/menu';
import {MatDialogModule} from '@angular/material/dialog';
import {DialogListComponent} from './component/dialog/dialog-list.component';
import {GroupIndexComponent} from './component/groups/index/group-index.component';
import {UserPolicyComponent} from './component/users/policy/user-policy.component';
import {PolicyService} from './service/policy.service';
import {GroupUserComponent} from './component/groups/user/group-user.component';
import {GroupPolicyComponent} from './component/groups/policy/group-policy.component';
import {GroupRoleComponent} from './component/groups/role/group-role.component';
import {GroupPermissionComponent} from './component/groups/permission/group-permission.component';
import {GroupService} from './service/group.service';
import {AppService} from './service/app.service';
import {ServiceNewComponent} from './component/services/new/service-new.component';

@NgModule({
  declarations: [
    AppComponent,
    DialogListComponent,
    LoginComponent,
    UserIndexComponent,
    UserPolicyComponent,
    GroupIndexComponent,
    GroupUserComponent,
    GroupPolicyComponent,
    GroupRoleComponent,
    GroupPermissionComponent,
    ServiceNewComponent,
  ],
  imports: [
    HttpClientModule,
    ToastrModule.forRoot(),
    BrowserModule,
    BrowserAnimationsModule,
    AppRoutingModule,
    MatToolbarModule,
    MatIconModule,
    MatSidenavModule,
    MatListModule,
    MatCardModule,
    MatFormFieldModule,
    MatSelectModule,
    MatInputModule,
    ReactiveFormsModule,
    MatButtonModule,
    MatSelectModule,
    MatTableModule,
    MatMenuModule,
    MatDialogModule,
    FormsModule,
    OverlayModule,
  ],
  providers: [
    CookieService,
    AppService,
    UserService,
    ServiceService,
    PolicyService,
    GroupService,
    LoginRequireGuard,
    LoggingInGuard,
  ],
  bootstrap: [AppComponent],
  schemas: [CUSTOM_ELEMENTS_SCHEMA],
  entryComponents: [MatSpinner],
})
export class AppModule {
}
