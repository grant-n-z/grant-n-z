import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { GroupPolicyComponent } from './group-policy.component';

describe('GroupPolicyComponent', () => {
  let component: GroupPolicyComponent;
  let fixture: ComponentFixture<GroupPolicyComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ GroupPolicyComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(GroupPolicyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
