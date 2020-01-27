import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { SideDrawerComponent } from './side-drawer.component';
import { ItemService } from 'src/app/services/item.service';

describe('SideDrawerComponent', () => {
  let component: SideDrawerComponent;
  let itemService: ItemService;
  let fixture: ComponentFixture<SideDrawerComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ SideDrawerComponent ],
      providers: [{provide:ItemService, useValue:itemService}]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(SideDrawerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
