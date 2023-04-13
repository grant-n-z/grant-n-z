import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupPermissionComponent } from './group-permission.component';

describe('GroupPermissionComponent', () => {
  let component: GroupPermissionComponent;
  let fixture: ComponentFixture<GroupPermissionComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GroupPermissionComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GroupPermissionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
