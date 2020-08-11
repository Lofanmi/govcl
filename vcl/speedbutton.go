
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TSpeedButton struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// 创建一个新的对象。
// 
// Create a new object.
func NewSpeedButton(owner IComponent) *TSpeedButton {
    s := new(TSpeedButton)
    s.instance = SpeedButton_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// 动态转换一个已存在的对象实例。
// 
// Dynamically convert an existing object instance.
func AsSpeedButton(obj interface{}) *TSpeedButton {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TSpeedButton{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// 新建一个对象来自已经存在的对象实例指针。
// 
// Create a new object from an existing object instance pointer.
// Deprecated: use AsSpeedButton.
func SpeedButtonFromInst(inst uintptr) *TSpeedButton {
    return AsSpeedButton(inst)
}

// 新建一个对象来自已经存在的对象实例。
// 
// Create a new object from an existing object instance.
// Deprecated: use AsSpeedButton.
func SpeedButtonFromObj(obj IObject) *TSpeedButton {
    return AsSpeedButton(obj)
}

// 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// 
// Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsSpeedButton.
func SpeedButtonFromUnsafePointer(ptr unsafe.Pointer) *TSpeedButton {
    return AsSpeedButton(ptr)
}

// -------------------------- Deprecated end --------------------------
// 释放对象。
// 
// Free object.
func (s *TSpeedButton) Free() {
    if s.instance != 0 {
        SpeedButton_Free(s.instance)
        s.instance, s.ptr = 0, nullptr
    }
}

// 返回对象实例指针。
// 
// Return object instance pointer.
func (s *TSpeedButton) Instance() uintptr {
    return s.instance
}

// 获取一个不安全的地址。
// 
// Get an unsafe address.
func (s *TSpeedButton) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// 检测地址是否为空。
// 
// Check if the address is empty.
func (s *TSpeedButton) IsValid() bool {
    return s.instance != 0
}

// 检测当前对象是否继承自目标对象。
// 
// Checks whether the current object is inherited from the target object.
func (s *TSpeedButton) Is() TIs {
    return TIs(s.instance)
}

// 动态转换当前对象为目标对象。
// 
// Dynamically convert the current object to the target object.
//func (s *TSpeedButton) As() TAs {
//    return TAs(s.instance)
//}

// 获取类信息指针。
// 
// Get class information pointer.
func TSpeedButtonClass() TClass {
    return SpeedButton_StaticClassType()
}

// 单击。
func (s *TSpeedButton) Click() {
    SpeedButton_Click(s.instance)
}

// 将控件置于最前。
//
// Bring the control to the front.
func (s *TSpeedButton) BringToFront() {
    SpeedButton_BringToFront(s.instance)
}

// 将客户端坐标转为绝对的屏幕坐标。
//
// Convert client coordinates to absolute screen coordinates.
func (s *TSpeedButton) ClientToScreen(Point TPoint) TPoint {
    return SpeedButton_ClientToScreen(s.instance, Point)
}

// 将客户端坐标转为父容器坐标。
//
// Convert client coordinates to parent container coordinates.
func (s *TSpeedButton) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return SpeedButton_ClientToParent(s.instance, Point , CheckPtr(AParent))
}

// 是否在拖拽中。
//
// Is it in the middle of dragging.
func (s *TSpeedButton) Dragging() bool {
    return SpeedButton_Dragging(s.instance)
}

// 是否有父容器。
//
// Is there a parent container.
func (s *TSpeedButton) HasParent() bool {
    return SpeedButton_HasParent(s.instance)
}

// 隐藏控件。
//
// Hidden control.
func (s *TSpeedButton) Hide() {
    SpeedButton_Hide(s.instance)
}

// 要求重绘。
//
// Redraw.
func (s *TSpeedButton) Invalidate() {
    SpeedButton_Invalidate(s.instance)
}

// 发送一个消息。
//
// Send a message.
func (s *TSpeedButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return SpeedButton_Perform(s.instance, Msg , WParam , LParam)
}

// 刷新控件。
//
// Refresh control.
func (s *TSpeedButton) Refresh() {
    SpeedButton_Refresh(s.instance)
}

// 重绘。
//
// Repaint.
func (s *TSpeedButton) Repaint() {
    SpeedButton_Repaint(s.instance)
}

// 将屏幕坐标转为客户端坐标。
//
// Convert screen coordinates to client coordinates.
func (s *TSpeedButton) ScreenToClient(Point TPoint) TPoint {
    return SpeedButton_ScreenToClient(s.instance, Point)
}

// 将父容器坐标转为客户端坐标。
//
// Convert parent container coordinates to client coordinates.
func (s *TSpeedButton) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return SpeedButton_ParentToClient(s.instance, Point , CheckPtr(AParent))
}

// 控件至于最后面。
//
// The control is placed at the end.
func (s *TSpeedButton) SendToBack() {
    SpeedButton_SendToBack(s.instance)
}

// 设置组件边界。
//
// Set component boundaries.
func (s *TSpeedButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    SpeedButton_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

// 显示控件。
//
// Show control.
func (s *TSpeedButton) Show() {
    SpeedButton_Show(s.instance)
}

// 控件更新。
//
// Update.
func (s *TSpeedButton) Update() {
    SpeedButton_Update(s.instance)
}

// 获取控件的字符，如果有。
//
// Get the characters of the control, if any.
func (s *TSpeedButton) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return SpeedButton_GetTextBuf(s.instance, Buffer , BufSize)
}

// 获取控件的字符长，如果有。
//
// Get the character length of the control, if any.
func (s *TSpeedButton) GetTextLen() int32 {
    return SpeedButton_GetTextLen(s.instance)
}

// 设置控件字符，如果有。
//
// Set control characters, if any.
func (s *TSpeedButton) SetTextBuf(Buffer string) {
    SpeedButton_SetTextBuf(s.instance, Buffer)
}

// 查找指定名称的组件。
//
// Find the component with the specified name.
func (s *TSpeedButton) FindComponent(AName string) *TComponent {
    return AsComponent(SpeedButton_FindComponent(s.instance, AName))
}

// 获取类名路径。
//
// Get the class name path.
func (s *TSpeedButton) GetNamePath() string {
    return SpeedButton_GetNamePath(s.instance)
}

// 复制一个对象，如果对象实现了此方法的话。
//
// Copy an object, if the object implements this method.
func (s *TSpeedButton) Assign(Source IObject) {
    SpeedButton_Assign(s.instance, CheckPtr(Source))
}

// 获取类的类型信息。
//
// Get class type information.
func (s *TSpeedButton) ClassType() TClass {
    return SpeedButton_ClassType(s.instance)
}

// 获取当前对象类名称。
//
// Get the current object class name.
func (s *TSpeedButton) ClassName() string {
    return SpeedButton_ClassName(s.instance)
}

// 获取当前对象实例大小。
//
// Get the current object instance size.
func (s *TSpeedButton) InstanceSize() int32 {
    return SpeedButton_InstanceSize(s.instance)
}

// 判断当前类是否继承自指定类。
//
// Determine whether the current class inherits from the specified class.
func (s *TSpeedButton) InheritsFrom(AClass TClass) bool {
    return SpeedButton_InheritsFrom(s.instance, AClass)
}

// 与一个对象进行比较。
//
// Compare with an object.
func (s *TSpeedButton) Equals(Obj IObject) bool {
    return SpeedButton_Equals(s.instance, CheckPtr(Obj))
}

// 获取类的哈希值。
//
// Get the hash value of the class.
func (s *TSpeedButton) GetHashCode() int32 {
    return SpeedButton_GetHashCode(s.instance)
}

// 文本类信息。
//
// Text information.
func (s *TSpeedButton) ToString() string {
    return SpeedButton_ToString(s.instance)
}

func (s *TSpeedButton) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    SpeedButton_AnchorToNeighbour(s.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (s *TSpeedButton) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    SpeedButton_AnchorParallel(s.instance, ASide , ASpace , CheckPtr(ASibling))
}

// 置于指定控件的横向中心。
func (s *TSpeedButton) AnchorHorizontalCenterTo(ASibling IControl) {
    SpeedButton_AnchorHorizontalCenterTo(s.instance, CheckPtr(ASibling))
}

// 置于指定控件的纵向中心。
func (s *TSpeedButton) AnchorVerticalCenterTo(ASibling IControl) {
    SpeedButton_AnchorVerticalCenterTo(s.instance, CheckPtr(ASibling))
}

func (s *TSpeedButton) AnchorSame(ASide TAnchorKind, ASibling IControl) {
    SpeedButton_AnchorSame(s.instance, ASide , CheckPtr(ASibling))
}

func (s *TSpeedButton) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    SpeedButton_AnchorAsAlign(s.instance, ATheAlign , ASpace)
}

func (s *TSpeedButton) AnchorClient(ASpace int32) {
    SpeedButton_AnchorClient(s.instance, ASpace)
}

// 获取图像在images中的索引。
func (s *TSpeedButton) ImageIndex() int32 {
    return SpeedButton_GetImageIndex(s.instance)
}

// 设置图像在images中的索引。
func (s *TSpeedButton) SetImageIndex(value int32) {
    SpeedButton_SetImageIndex(s.instance, value)
}

// 获取图标索引列表对象。
func (s *TSpeedButton) Images() *TImageList {
    return AsImageList(SpeedButton_GetImages(s.instance))
}

// 设置图标索引列表对象。
func (s *TSpeedButton) SetImages(value IComponent) {
    SpeedButton_SetImages(s.instance, CheckPtr(value))
}

func (s *TSpeedButton) ImageWidth() int32 {
    return SpeedButton_GetImageWidth(s.instance)
}

func (s *TSpeedButton) SetImageWidth(value int32) {
    SpeedButton_SetImageWidth(s.instance, value)
}

func (s *TSpeedButton) ShowCaption() bool {
    return SpeedButton_GetShowCaption(s.instance)
}

func (s *TSpeedButton) SetShowCaption(value bool) {
    SpeedButton_SetShowCaption(s.instance, value)
}

func (s *TSpeedButton) Action() *TAction {
    return AsAction(SpeedButton_GetAction(s.instance))
}

func (s *TSpeedButton) SetAction(value IComponent) {
    SpeedButton_SetAction(s.instance, CheckPtr(value))
}

// 获取控件自动调整。
//
// Get Control automatically adjusts.
func (s *TSpeedButton) Align() TAlign {
    return SpeedButton_GetAlign(s.instance)
}

// 设置控件自动调整。
//
// Set Control automatically adjusts.
func (s *TSpeedButton) SetAlign(value TAlign) {
    SpeedButton_SetAlign(s.instance, value)
}

func (s *TSpeedButton) AllowAllUp() bool {
    return SpeedButton_GetAllowAllUp(s.instance)
}

func (s *TSpeedButton) SetAllowAllUp(value bool) {
    SpeedButton_SetAllowAllUp(s.instance, value)
}

// 获取四个角位置的锚点。
func (s *TSpeedButton) Anchors() TAnchors {
    return SpeedButton_GetAnchors(s.instance)
}

// 设置四个角位置的锚点。
func (s *TSpeedButton) SetAnchors(value TAnchors) {
    SpeedButton_SetAnchors(s.instance, value)
}

func (s *TSpeedButton) BiDiMode() TBiDiMode {
    return SpeedButton_GetBiDiMode(s.instance)
}

func (s *TSpeedButton) SetBiDiMode(value TBiDiMode) {
    SpeedButton_SetBiDiMode(s.instance, value)
}

// 获取约束控件大小。
func (s *TSpeedButton) Constraints() *TSizeConstraints {
    return AsSizeConstraints(SpeedButton_GetConstraints(s.instance))
}

// 设置约束控件大小。
func (s *TSpeedButton) SetConstraints(value *TSizeConstraints) {
    SpeedButton_SetConstraints(s.instance, CheckPtr(value))
}

// 获取团组索引。
func (s *TSpeedButton) GroupIndex() int32 {
    return SpeedButton_GetGroupIndex(s.instance)
}

// 设置团组索引。
func (s *TSpeedButton) SetGroupIndex(value int32) {
    SpeedButton_SetGroupIndex(s.instance, value)
}

func (s *TSpeedButton) Down() bool {
    return SpeedButton_GetDown(s.instance)
}

func (s *TSpeedButton) SetDown(value bool) {
    SpeedButton_SetDown(s.instance, value)
}

// 获取控件标题。
//
// Get the control title.
func (s *TSpeedButton) Caption() string {
    return SpeedButton_GetCaption(s.instance)
}

// 设置控件标题。
//
// Set the control title.
func (s *TSpeedButton) SetCaption(value string) {
    SpeedButton_SetCaption(s.instance, value)
}

// 获取控件启用。
//
// Get the control enabled.
func (s *TSpeedButton) Enabled() bool {
    return SpeedButton_GetEnabled(s.instance)
}

// 设置控件启用。
//
// Set the control enabled.
func (s *TSpeedButton) SetEnabled(value bool) {
    SpeedButton_SetEnabled(s.instance, value)
}

// 获取平面样式。
func (s *TSpeedButton) Flat() bool {
    return SpeedButton_GetFlat(s.instance)
}

// 设置平面样式。
func (s *TSpeedButton) SetFlat(value bool) {
    SpeedButton_SetFlat(s.instance, value)
}

// 获取字体。
//
// Get Font.
func (s *TSpeedButton) Font() *TFont {
    return AsFont(SpeedButton_GetFont(s.instance))
}

// 设置字体。
//
// Set Font.
func (s *TSpeedButton) SetFont(value *TFont) {
    SpeedButton_SetFont(s.instance, CheckPtr(value))
}

func (s *TSpeedButton) Glyph() *TBitmap {
    return AsBitmap(SpeedButton_GetGlyph(s.instance))
}

func (s *TSpeedButton) SetGlyph(value *TBitmap) {
    SpeedButton_SetGlyph(s.instance, CheckPtr(value))
}

func (s *TSpeedButton) Layout() TButtonLayout {
    return SpeedButton_GetLayout(s.instance)
}

func (s *TSpeedButton) SetLayout(value TButtonLayout) {
    SpeedButton_SetLayout(s.instance, value)
}

func (s *TSpeedButton) NumGlyphs() TNumGlyphs {
    return SpeedButton_GetNumGlyphs(s.instance)
}

func (s *TSpeedButton) SetNumGlyphs(value TNumGlyphs) {
    SpeedButton_SetNumGlyphs(s.instance, value)
}

// 获取使用父容器字体。
//
// Get Parent container font.
func (s *TSpeedButton) ParentFont() bool {
    return SpeedButton_GetParentFont(s.instance)
}

// 设置使用父容器字体。
//
// Set Parent container font.
func (s *TSpeedButton) SetParentFont(value bool) {
    SpeedButton_SetParentFont(s.instance, value)
}

// 获取以父容器的ShowHint属性为准。
func (s *TSpeedButton) ParentShowHint() bool {
    return SpeedButton_GetParentShowHint(s.instance)
}

// 设置以父容器的ShowHint属性为准。
func (s *TSpeedButton) SetParentShowHint(value bool) {
    SpeedButton_SetParentShowHint(s.instance, value)
}

// 获取右键菜单。
//
// Get Right click menu.
func (s *TSpeedButton) PopupMenu() *TPopupMenu {
    return AsPopupMenu(SpeedButton_GetPopupMenu(s.instance))
}

// 设置右键菜单。
//
// Set Right click menu.
func (s *TSpeedButton) SetPopupMenu(value IComponent) {
    SpeedButton_SetPopupMenu(s.instance, CheckPtr(value))
}

// 获取显示鼠标悬停提示。
//
// Get Show mouseover tips.
func (s *TSpeedButton) ShowHint() bool {
    return SpeedButton_GetShowHint(s.instance)
}

// 设置显示鼠标悬停提示。
//
// Set Show mouseover tips.
func (s *TSpeedButton) SetShowHint(value bool) {
    SpeedButton_SetShowHint(s.instance, value)
}

func (s *TSpeedButton) Spacing() int32 {
    return SpeedButton_GetSpacing(s.instance)
}

func (s *TSpeedButton) SetSpacing(value int32) {
    SpeedButton_SetSpacing(s.instance, value)
}

// 获取透明。
//
// Get transparent.
func (s *TSpeedButton) Transparent() bool {
    return SpeedButton_GetTransparent(s.instance)
}

// 设置透明。
//
// Set transparent.
func (s *TSpeedButton) SetTransparent(value bool) {
    SpeedButton_SetTransparent(s.instance, value)
}

// 获取控件可视。
//
// Get the control visible.
func (s *TSpeedButton) Visible() bool {
    return SpeedButton_GetVisible(s.instance)
}

// 设置控件可视。
//
// Set the control visible.
func (s *TSpeedButton) SetVisible(value bool) {
    SpeedButton_SetVisible(s.instance, value)
}

// 设置控件单击事件。
//
// Set control click event.
func (s *TSpeedButton) SetOnClick(fn TNotifyEvent) {
    SpeedButton_SetOnClick(s.instance, fn)
}

// 设置双击事件。
func (s *TSpeedButton) SetOnDblClick(fn TNotifyEvent) {
    SpeedButton_SetOnDblClick(s.instance, fn)
}

// 设置鼠标按下事件。
//
// Set Mouse down event.
func (s *TSpeedButton) SetOnMouseDown(fn TMouseEvent) {
    SpeedButton_SetOnMouseDown(s.instance, fn)
}

// 设置鼠标进入事件。
//
// Set Mouse entry event.
func (s *TSpeedButton) SetOnMouseEnter(fn TNotifyEvent) {
    SpeedButton_SetOnMouseEnter(s.instance, fn)
}

// 设置鼠标离开事件。
//
// Set Mouse leave event.
func (s *TSpeedButton) SetOnMouseLeave(fn TNotifyEvent) {
    SpeedButton_SetOnMouseLeave(s.instance, fn)
}

// 设置鼠标移动事件。
func (s *TSpeedButton) SetOnMouseMove(fn TMouseMoveEvent) {
    SpeedButton_SetOnMouseMove(s.instance, fn)
}

// 设置鼠标抬起事件。
//
// Set Mouse lift event.
func (s *TSpeedButton) SetOnMouseUp(fn TMouseEvent) {
    SpeedButton_SetOnMouseUp(s.instance, fn)
}

func (s *TSpeedButton) BoundsRect() TRect {
    return SpeedButton_GetBoundsRect(s.instance)
}

func (s *TSpeedButton) SetBoundsRect(value TRect) {
    SpeedButton_SetBoundsRect(s.instance, value)
}

// 获取客户区高度。
//
// Get client height.
func (s *TSpeedButton) ClientHeight() int32 {
    return SpeedButton_GetClientHeight(s.instance)
}

// 设置客户区高度。
//
// Set client height.
func (s *TSpeedButton) SetClientHeight(value int32) {
    SpeedButton_SetClientHeight(s.instance, value)
}

func (s *TSpeedButton) ClientOrigin() TPoint {
    return SpeedButton_GetClientOrigin(s.instance)
}

// 获取客户区矩形。
//
// Get client rectangle.
func (s *TSpeedButton) ClientRect() TRect {
    return SpeedButton_GetClientRect(s.instance)
}

// 获取客户区宽度。
//
// Get client width.
func (s *TSpeedButton) ClientWidth() int32 {
    return SpeedButton_GetClientWidth(s.instance)
}

// 设置客户区宽度。
//
// Set client width.
func (s *TSpeedButton) SetClientWidth(value int32) {
    SpeedButton_SetClientWidth(s.instance, value)
}

// 获取控件状态。
//
// Get control state.
func (s *TSpeedButton) ControlState() TControlState {
    return SpeedButton_GetControlState(s.instance)
}

// 设置控件状态。
//
// Set control state.
func (s *TSpeedButton) SetControlState(value TControlState) {
    SpeedButton_SetControlState(s.instance, value)
}

// 获取控件样式。
//
// Get control style.
func (s *TSpeedButton) ControlStyle() TControlStyle {
    return SpeedButton_GetControlStyle(s.instance)
}

// 设置控件样式。
//
// Set control style.
func (s *TSpeedButton) SetControlStyle(value TControlStyle) {
    SpeedButton_SetControlStyle(s.instance, value)
}

func (s *TSpeedButton) Floating() bool {
    return SpeedButton_GetFloating(s.instance)
}

// 获取控件父容器。
//
// Get control parent container.
func (s *TSpeedButton) Parent() *TWinControl {
    return AsWinControl(SpeedButton_GetParent(s.instance))
}

// 设置控件父容器。
//
// Set control parent container.
func (s *TSpeedButton) SetParent(value IWinControl) {
    SpeedButton_SetParent(s.instance, CheckPtr(value))
}

// 获取左边位置。
//
// Get Left position.
func (s *TSpeedButton) Left() int32 {
    return SpeedButton_GetLeft(s.instance)
}

// 设置左边位置。
//
// Set Left position.
func (s *TSpeedButton) SetLeft(value int32) {
    SpeedButton_SetLeft(s.instance, value)
}

// 获取顶边位置。
//
// Get Top position.
func (s *TSpeedButton) Top() int32 {
    return SpeedButton_GetTop(s.instance)
}

// 设置顶边位置。
//
// Set Top position.
func (s *TSpeedButton) SetTop(value int32) {
    SpeedButton_SetTop(s.instance, value)
}

// 获取宽度。
//
// Get width.
func (s *TSpeedButton) Width() int32 {
    return SpeedButton_GetWidth(s.instance)
}

// 设置宽度。
//
// Set width.
func (s *TSpeedButton) SetWidth(value int32) {
    SpeedButton_SetWidth(s.instance, value)
}

// 获取高度。
//
// Get height.
func (s *TSpeedButton) Height() int32 {
    return SpeedButton_GetHeight(s.instance)
}

// 设置高度。
//
// Set height.
func (s *TSpeedButton) SetHeight(value int32) {
    SpeedButton_SetHeight(s.instance, value)
}

// 获取控件光标。
//
// Get control cursor.
func (s *TSpeedButton) Cursor() TCursor {
    return SpeedButton_GetCursor(s.instance)
}

// 设置控件光标。
//
// Set control cursor.
func (s *TSpeedButton) SetCursor(value TCursor) {
    SpeedButton_SetCursor(s.instance, value)
}

// 获取组件鼠标悬停提示。
//
// Get component mouse hints.
func (s *TSpeedButton) Hint() string {
    return SpeedButton_GetHint(s.instance)
}

// 设置组件鼠标悬停提示。
//
// Set component mouse hints.
func (s *TSpeedButton) SetHint(value string) {
    SpeedButton_SetHint(s.instance, value)
}

// 获取组件总数。
//
// Get the total number of components.
func (s *TSpeedButton) ComponentCount() int32 {
    return SpeedButton_GetComponentCount(s.instance)
}

// 获取组件索引。
//
// Get component index.
func (s *TSpeedButton) ComponentIndex() int32 {
    return SpeedButton_GetComponentIndex(s.instance)
}

// 设置组件索引。
//
// Set component index.
func (s *TSpeedButton) SetComponentIndex(value int32) {
    SpeedButton_SetComponentIndex(s.instance, value)
}

// 获取组件所有者。
//
// Get component owner.
func (s *TSpeedButton) Owner() *TComponent {
    return AsComponent(SpeedButton_GetOwner(s.instance))
}

// 获取组件名称。
//
// Get the component name.
func (s *TSpeedButton) Name() string {
    return SpeedButton_GetName(s.instance)
}

// 设置组件名称。
//
// Set the component name.
func (s *TSpeedButton) SetName(value string) {
    SpeedButton_SetName(s.instance, value)
}

// 获取对象标记。
//
// Get the control tag.
func (s *TSpeedButton) Tag() int {
    return SpeedButton_GetTag(s.instance)
}

// 设置对象标记。
//
// Set the control tag.
func (s *TSpeedButton) SetTag(value int) {
    SpeedButton_SetTag(s.instance, value)
}

// 获取左边锚点。
func (s *TSpeedButton) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(SpeedButton_GetAnchorSideLeft(s.instance))
}

// 设置左边锚点。
func (s *TSpeedButton) SetAnchorSideLeft(value *TAnchorSide) {
    SpeedButton_SetAnchorSideLeft(s.instance, CheckPtr(value))
}

// 获取顶边锚点。
func (s *TSpeedButton) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(SpeedButton_GetAnchorSideTop(s.instance))
}

// 设置顶边锚点。
func (s *TSpeedButton) SetAnchorSideTop(value *TAnchorSide) {
    SpeedButton_SetAnchorSideTop(s.instance, CheckPtr(value))
}

// 获取右边锚点。
func (s *TSpeedButton) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(SpeedButton_GetAnchorSideRight(s.instance))
}

// 设置右边锚点。
func (s *TSpeedButton) SetAnchorSideRight(value *TAnchorSide) {
    SpeedButton_SetAnchorSideRight(s.instance, CheckPtr(value))
}

// 获取底边锚点。
func (s *TSpeedButton) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(SpeedButton_GetAnchorSideBottom(s.instance))
}

// 设置底边锚点。
func (s *TSpeedButton) SetAnchorSideBottom(value *TAnchorSide) {
    SpeedButton_SetAnchorSideBottom(s.instance, CheckPtr(value))
}

// 获取边框间距。
func (s *TSpeedButton) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(SpeedButton_GetBorderSpacing(s.instance))
}

// 设置边框间距。
func (s *TSpeedButton) SetBorderSpacing(value *TControlBorderSpacing) {
    SpeedButton_SetBorderSpacing(s.instance, CheckPtr(value))
}

// 获取指定索引组件。
//
// Get the specified index component.
func (s *TSpeedButton) Components(AIndex int32) *TComponent {
    return AsComponent(SpeedButton_GetComponents(s.instance, AIndex))
}

// 获取锚侧面。
func (s *TSpeedButton) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(SpeedButton_GetAnchorSide(s.instance, AKind))
}

