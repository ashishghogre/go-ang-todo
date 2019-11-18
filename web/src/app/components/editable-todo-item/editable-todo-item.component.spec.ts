import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EditableTodoItemComponent } from './editable-todo-item.component';

describe('EditableTodoItemComponent', () => {
  let component: EditableTodoItemComponent;
  let fixture: ComponentFixture<EditableTodoItemComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EditableTodoItemComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EditableTodoItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
